package stack

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (p *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (p *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
