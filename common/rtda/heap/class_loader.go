package heap

import (
	"fmt"
	"gojvm/common/classfile"
	"gojvm/common/classpath"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (cl *ClassLoader) LoadClass(name string) *Class {
	if class, ok := cl.classMap[name]; ok {
		return class
	}
	return cl.loadNonArrayClass(name)
}

// loadNonArrayClass 加载非数组类
// 1. 从class文件读取类数据
// 2. 解析class文件，生成虚拟机可以使用的类数据,并放入方法区
// 3. 进行链接
func (cl *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := cl.readClass(name)
	class := cl.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (cl *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := cl.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (cl *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = cl
	resolveSuperClass(class)
	resolveInterfaces(class)
	cl.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

// 解析超类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 解析接口
func resolveInterfaces(class *Class) {
	if len(class.interfaceNames) > 0 {
		interfaces := make([]*Class, len(class.interfaceNames))
		for i, interfaceName := range class.interfaceNames {
			interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

// 类的验证，为了确保安全性，Java虚拟机规范要求在执行类的任何代码之前，对类进行严格的验证
func verify(class *Class) {
	// todo
}

// 准备阶段，给类变量空间分配初始值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// calcInstanceFieldSlotIds 用于计算实例字段的个数，同时给它们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if field.IsStatic() {
			continue // 静态的就不计算
		}
		field.slotId = slotId
		slotId++
		if field.IsLongOrDouble() {
			slotId++
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstantValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
