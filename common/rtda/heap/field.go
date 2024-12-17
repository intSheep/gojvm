package heap

import "gojvm/common/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (f *Field) SlotID() uint {
	return f.slotId
}

func (f *Field) Class() *Class {
	return f.class
}

func (f *Field) IsStatic() bool {
	return f.accessFlags&ACC_STATIC != 0
}

func (f *Field) IsPrivate() bool {
	return 0 != f.accessFlags&ACC_PRIVATE
}

func (f *Field) IsFinal() bool {
	return 0 != f.accessFlags&ACC_FINAL
}

func (f *Field) IsLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}

func (f *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		f.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (f *Field) IsPublic() bool {
	return 0 != f.accessFlags&ACC_PUBLIC
}

func (f *Field) IsProtected() bool {
	return 0 != f.accessFlags&ACC_PROTECTED
}

func (f *Field) IsAccessibleTo(c *Class) bool {
	if f.IsPublic() {
		return true
	}
	d := f.class
	if f.IsProtected() {
		// protected字段只能被同一个包下的类访问，或者子类访问（class-c是class-d的子类）
		return c == d || d.isSubClassOf(c) || d.getPackageName() == c.getPackageName()
	}
	if !f.IsPrivate() {
		return d.getPackageName() == c.getPackageName()
	}
	return d == c
}

func (f *Field) ConstantValueIndex() uint {
	return f.constValueIndex
}

func (f *Field) SlotId() uint {
	return f.slotId
}
