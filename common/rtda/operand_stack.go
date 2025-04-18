package rtda

import (
	"gojvm/common/rtda/heap"
	"math"
)

// OperandStack 操作数栈
// 用于存储字节码指令操作的数据和中间结果
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].num = int32(bits)
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	bits := uint32(o.slots[o.size].num)
	return math.Float32frombits(bits)
}

func (o *OperandStack) PushLong(val int64) {
	o.slots[o.size].num = int32(val)
	o.slots[o.size+1].num = int32(val >> 32)
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (o *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	o.PushLong(int64(bits))
}

func (o *OperandStack) PopDouble() float64 {
	bits := uint64(o.PopLong())
	return math.Float64frombits(bits)
}

func (o *OperandStack) PushRef(ref *heap.Object) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *heap.Object {
	o.size--
	ref := o.slots[o.size].ref
	o.slots[o.size].ref = nil
	return ref
}

func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	return o.slots[o.size]
}
