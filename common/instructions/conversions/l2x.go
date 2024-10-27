package conversions

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type L2D struct {
	base.NoOperandsInstruction
}

func (l *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	lVal := stack.PopLong()
	dVal := float64(lVal)
	stack.PushDouble(dVal)
}

type L2F struct {
	base.NoOperandsInstruction
}

func (l *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	lVal := stack.PopLong()
	fVal := float32(lVal)
	stack.PushFloat(fVal)
}

type L2I struct {
	base.NoOperandsInstruction
}

func (l *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	lVal := stack.PopLong()
	iVal := int32(lVal)
	stack.PushInt(iVal)
}
