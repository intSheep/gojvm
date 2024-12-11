package heap

import "gojvm/common/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRef(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {
		f.resolveFieldRef()
	}
	return f.field
}

func (f *FieldRef) resolveFieldRef() {
	d := f.cp.class
	c := f.ResolvedClass()
	field := lookupField(c, f.name, f.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	f.field = field
}

// lookupField 在类中查找字段
// 先在当前类中查找，然后在接口中查找，最后在父类中查找
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
