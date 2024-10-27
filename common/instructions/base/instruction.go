package base

import (
	"gojvm/common/rtda"
)

type Instructions interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (n *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
}

func (n *NoOperandsInstruction) Execute(frame *rtda.Frame) {
}

// BranchInstruction 跳转指令，其中Offset为跳转偏移量
type BranchInstruction struct {
	Offset int
}

func (b *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	b.Offset = int(reader.ReadInt16())
}

// Index8Instruction 操作数为8位的指令
type Index8Instruction struct {
	Index uint
}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint16())
}
