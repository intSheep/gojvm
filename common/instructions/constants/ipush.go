package constants

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type BIPUSH struct {
	val int8
}

func (b *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	b.val = reader.ReadInt8()
}

func (b *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(b.val))
}

type SIPUSH struct {
	val int16
}

func (s *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	s.val = reader.ReadInt16()
}
func (s *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(s.val)
	frame.OperandStack().PushInt(i)
}
