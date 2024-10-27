package comparisons

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type ICMP struct {
	base.NoOperandsInstruction
}

func (i *ICMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
