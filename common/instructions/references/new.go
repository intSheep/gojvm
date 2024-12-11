package references

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
	"gojvm/common/rtda/heap"
)

type New struct {
	base.Index16Instruction
}

// Execute 实例化一个类
func (n *New) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(n.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// 接口类和抽象类不能实例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
