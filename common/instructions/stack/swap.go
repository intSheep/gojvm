package stack

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (s *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
