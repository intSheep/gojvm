package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (c *ExceptionsAttribute) readInfo(reader *ClassReader) {
	c.exceptionIndexTable = reader.readUint16s()
}

func (c *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return c.exceptionIndexTable
}
