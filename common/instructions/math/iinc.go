package math

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

// IINC 给局部变量表的int常量增加常量值
// 其中索引和常量都是由操作数提供
type IINC struct {
	Index uint
	Const int32
}

func (i *IINC) FetchOperands(reader *base.BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
	i.Const = int32(reader.ReadInt8())
}

func (i *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(i.Index)
	val += i.Const
	localVars.SetInt(i.Index, val)
}
