package references

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
	"gojvm/common/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

// Execute 给类的静态变量赋值
func (p *PUT_STATIC) Execute(frame *rtda.Frame) {
	// 拿出常量池
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	// 找到字段符号引用
	fieldRef := cp.GetConstant(p.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果是final字段，只能在类初始化方法中赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	// 通过field的描述符决定如何给静态变量赋值
	descriptor := field.Descriptor()
	slotId := field.SlotID()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
