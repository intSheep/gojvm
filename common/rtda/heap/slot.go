package heap

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(num uint) Slots {
	return make(Slots, num)
}

func (s Slots) SetInt(index uint, val int32) {
	s[index].num = val
}

func (s Slots) SetFloat(index uint, val float32) {
	bits := int32(val)
	s[index].num = bits
}

func (s Slots) SetLong(index uint, val int64) {
	s[index].num = int32(val)
	s[index+1].num = int32(val >> 32)
}

func (s Slots) SetDouble(index uint, val float64) {
	bits := int64(val)
	s.SetLong(index, bits)
}

func (s Slots) SetRef(index uint, ref *Object) {
	s[index].ref = ref
}

func (s Slots) GetInt(index uint) int32 {
	return s[index].num
}

func (s Slots) GetFloat(index uint) float32 {
	bits := uint32(s[index].num)
	return float32(bits)
}

func (s Slots) GetLong(index uint) int64 {
	low := uint32(s[index].num)
	high := uint32(s[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (s Slots) GetDouble(index uint) float64 {
	bits := uint64(s.GetLong(index))
	return float64(bits)
}

func (s Slots) GetRef(index uint) *Object {
	return s[index].ref
}
