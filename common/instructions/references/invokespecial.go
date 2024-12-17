package references

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (i *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
