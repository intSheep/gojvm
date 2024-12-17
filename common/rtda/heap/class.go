package heap

import (
	"gojvm/common/classfile"
	"strings"
)

const (
	ACC_PUBLIC       = 0x0001 // class field method
	ACC_PRIVATE      = 0x0002 //         field method
	ACC_PROTECTED    = 0x0004 //         field method
	ACC_STATIC       = 0x0008 //         field method
	ACC_FINAL        = 0x0010 // class field method
	ACC_SUPER        = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 //                 method
	ACC_VOLATILE     = 0x0040 //         field
	ACC_BRIDGE       = 0x0040 //                 method
	ACC_TRANSIENT    = 0x0080 //         field
	ACC_VARARGS      = 0x0080 //                 method
	ACC_NATIVE       = 0x0100 //                 method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class         method
	ACC_STRICT       = 0x0800 //                 method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)

type Class struct {
	accessFlags       uint16
	name              string //全限定名
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) NewObject() *Object {
	return newObject(c)
}

func (c *Class) IsPublic() bool {
	return 0 != c.accessFlags&ACC_PUBLIC
}

func (c *Class) IsAbstract() bool {
	return 0 != c.accessFlags&ACC_ABSTRACT
}

func (c *Class) IsInterface() bool {
	return 0 != c.accessFlags&ACC_INTERFACE
}

// IsAccessibleTo 判断c类是否可以被other类访问
// 如果c类是public，或者c类和other类在同一个包下，那么c类就可以被访问
func (c *Class) IsAccessibleTo(other *Class) bool {
	return c.IsPublic() || c.getPackageName() == other.getPackageName()
}

func (c *Class) getPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

func (c *Class) isSubClassOf(d *Class) bool {
	if c == d {
		return true
	}
	if c.superClass != nil {
		return c.superClass.isSubClassOf(d)
	}
	return false
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}
