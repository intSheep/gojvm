package rtda

// Thread 线程
// 线程是执行代码的载体，而栈帧则是线程执行方法时的基本工作单位，线程通过其虚拟机栈来管理和执行一系列的栈帧。
type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (t *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(t, maxLocals, maxStack)
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.peek()
}
