package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func NewBytecodeReader() *BytecodeReader {
	return &BytecodeReader{}
}

func (br *BytecodeReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}
func (br *BytecodeReader) ReadUint8() uint8 {
	i := br.code[br.pc]
	br.pc++
	return i
}
func (br *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(br.ReadUint8())
	byte2 := uint16(br.ReadUint8())
	return (byte1 << 8) | byte2
}

func (br *BytecodeReader) ReadInt8() int8 {
	return int8(br.ReadUint8())
}

func (br *BytecodeReader) ReadInt16() int16 {
	return int16(br.ReadUint16())
}

func (br *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(br.ReadUint8())
	byte2 := int32(br.ReadUint8())
	byte3 := int32(br.ReadUint8())
	byte4 := int32(br.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (br *BytecodeReader) ReadInt32s(n int32) []int32 {
	s := make([]int32, n)
	for i := range s {
		s[i] = br.ReadInt32()
	}
	return s
}

func (br *BytecodeReader) SkipPadding() {
	for br.pc%4 != 0 {
		br.ReadUint8()
	}
}

func (br *BytecodeReader) PC() int {
	return br.pc
}
