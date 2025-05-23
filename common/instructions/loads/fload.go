package loads

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

func (f *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, f.Index)
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
