package loads

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type ILOAD struct {
	base.Index8Instruction
}

func (i *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, i.Index)
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
