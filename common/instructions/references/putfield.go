package references

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
	"gojvm/common/rtda/heap"
)

type PUT_FIELD struct {
	base.Index16Instruction
}

func (p *PUT_FIELD) Execute(frame *rtda.Frame) {
	method := frame.Method()
	currentClass := method.Class()
	cp := currentClass.ConstantPool()

	fieldRef := cp.GetConstant(p.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != field.Class() || method.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	slotId := field.SlotID()
	descriptor := field.Descriptor()
	stack := frame.OperandStack()

	// 从操作数栈中弹出字段值，然后弹出对象引用
	// 对象引用就是实例化对象
	// 如果对象引用为空，就会报空指针错误
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	}
}
