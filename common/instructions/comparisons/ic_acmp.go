package comparisons

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

// 判断引用是否一样
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (i *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (i *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, i.Offset)
	}
}
