package constants

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

// NOP Do Nothing
type NOP struct {
	base.NoOperandsInstruction
}

func (n *NOP) Execute(frame *rtda.Frame) {

}
