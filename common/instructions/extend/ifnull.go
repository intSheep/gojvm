package extend

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (i *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (i *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, i.Offset)
	}
}
