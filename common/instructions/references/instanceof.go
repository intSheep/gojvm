package references

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
	"gojvm/common/rtda/heap"
)

// INSTANCE_OF 判断对象是否是某个类的实例
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (i *INSTANCE_OF) Execute(frame *rtda.Frame) {
	// instance-of需要两个操作数
	// object instanceof class
	stack := frame.OperandStack()
	ref := stack.PopRef() //第一个操作数从操作数栈中弹出对象引用:object
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(i.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass() //第二个操作数是一个类符号引用:class
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
