package rtda

// Frame 栈帧
// 栈帧是用于存储数据和部分结果以及执行动态链接、返回值和分发异常的地方。
// 每个线程在执行一个方法时，都会为该方法创建一个新的栈帧并将其推入到线程的虚拟机栈中。
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int //当方法发生调用或者异常处理的时候，需要设置这个，指向下一条要执行的指令
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		thread:       thread,
	}
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) SetNextPC(next int) {
	f.nextPC = next
}

func (f *Frame) NextPC() int {
	return f.nextPC
}
