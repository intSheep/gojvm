package control

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type TABLE_SWITCH struct {
	// 默认情况下执行跳转所需的字节码偏移量
	defaultOffset int32
	// low和high记录case的取值范围
	low  int32
	high int32
	// jumpOffsets是一个索引表，里面存放high-low+1个int值，
	// 对应各种case情况下，执行跳转所需的字节码偏移量
	jumpOffsets []int32
}

func (t *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// tableswitch指令操作码的后面有0～3字节的padding，以保证defaultOffset在字节码中的地址是4的倍数
	reader.SkipPadding()
	t.defaultOffset = reader.ReadInt32()
	t.low = reader.ReadInt32()
	t.high = reader.ReadInt32()
	jumpOffsetsCount := t.high - t.low + 1
	t.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (t *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= t.low && index <= t.high {
		offset = int(t.jumpOffsets[index-t.low])
	} else {
		offset = int(t.defaultOffset)
	}
	base.Branch(frame, offset)
}
