package heap

import (
	"fmt"
	"gojvm/common/classfile"
)

type Constant interface{}

// ConstantPool 常量池，解析class文件也有一个常量池，但是那个是存储静态信息的
// 这个常量池是存储运行时数据的

type ConstantPool struct {
	class     *Class
	constants []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	consts := make([]Constant, len(cfCp))
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < len(cfCp); i++ {
		cpInfo := cfCp[i]
		switch inst := cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			consts[i] = inst.Value()
		case *classfile.ConstantFloatInfo:
			consts[i] = inst.Value()
		case *classfile.ConstantLongInfo:
			consts[i] = inst.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			consts[i] = inst.Value()
			i++
		case *classfile.ConstantStringInfo:
			consts[i] = inst.String()
		case *classfile.ConstantClassInfo:
			consts[i] = newClassRef(rtCp, inst)
		case *classfile.ConstantFieldRefInfo:
			consts[i] = newFieldRef(rtCp, inst)
		case *classfile.ConstantMethodRefInfo:
			consts[i] = newMethodRef(rtCp, inst)
		case *classfile.ConstantInterfaceMethodRefInfo:
			consts[i] = newInterfaceMethodRef(rtCp, inst)
		default:
		}
	}
	return rtCp
}

func (cp *ConstantPool) GetConstant(index uint) Constant {
	if c := cp.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
