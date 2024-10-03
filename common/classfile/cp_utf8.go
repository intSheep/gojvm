package classfile

type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	c.str = string(bytes)
}
