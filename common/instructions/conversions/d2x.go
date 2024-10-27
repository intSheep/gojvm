package conversions

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (d *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	f := float32(v)
	stack.PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (d *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	i := int32(v)
	stack.PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	l := int64(v)
	stack.PushLong(l)
}
