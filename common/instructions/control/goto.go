package control

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (g *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.Offset)
}
