package references

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
	"gojvm/common/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

// Execute 获取类的静态变量
func (g *GET_STATIC) Execute(frame *rtda.Frame) {
	// 拿出常量池
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	// 找到字段符号引用
	fieldRef := cp.GetConstant(g.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotID()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
