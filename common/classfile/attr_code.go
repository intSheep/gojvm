package classfile

type CodeAttribute struct {
	cp            ConstantPool
	maxStack      uint16
	maxLocals     uint16
	code          []byte
	exceptionable []*ExceptionTableEntry
	attributes    []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	c.maxStack = reader.readUint16()
	c.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	c.code = reader.readBytes(codeLength)
	c.exceptionable = readExceptionTable(reader)
	c.attributes = readAttributes(reader, c.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	expceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range expceptionTable {
		expceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return expceptionTable
}

func (c *CodeAttribute) MaxStack() uint {
	return uint(c.maxStack)
}

func (c *CodeAttribute) MaxLocals() uint {
	return uint(c.maxLocals)
}

func (c *CodeAttribute) Code() []byte {
	return c.code
}
