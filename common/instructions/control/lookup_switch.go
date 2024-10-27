package control

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32   // case值的数量
	matchOffsets  []int32 // 存储成对的case值，偶数是case的值，奇数是case对应偏移量
}

func (l *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	l.defaultOffset = reader.ReadInt32()
	l.npairs = reader.ReadInt32()
	l.matchOffsets = reader.ReadInt32s(l.npairs * 2)
}

func (l *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	key := stack.PopInt()
	for i := int32(0); i < l.npairs*2; i += 2 {
		if l.matchOffsets[i] == key {
			offset := int(l.matchOffsets[i+1])
			base.Branch(frame, offset)
			return
		}
	}
	base.Branch(frame, int(l.defaultOffset))
}
