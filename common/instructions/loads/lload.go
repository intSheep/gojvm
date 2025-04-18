package loads

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type LLOAD struct {
	base.Index8Instruction
}

func (l *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, l.Index)
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
