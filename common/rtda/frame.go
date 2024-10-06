package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals, maxStack int) *Frame {
	return &Frame{
		localVars:    nil,
		operandStack: nil,
	}
}
