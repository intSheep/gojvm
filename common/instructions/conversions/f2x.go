package conversions

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type F2I struct {
	base.NoOperandsInstruction
}

func (f *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	fVal := stack.PopFloat()
	iVal := int32(fVal)
	stack.PushInt(iVal)
}

type F2L struct {
	base.NoOperandsInstruction
}

func (f *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	fVal := stack.PopFloat()
	lVal := int64(fVal)
	stack.PushLong(lVal)
}

type F2D struct {
	base.NoOperandsInstruction
}

func (f *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	fVal := stack.PopFloat()
	dVal := float64(fVal)
	stack.PushDouble(dVal)
}
