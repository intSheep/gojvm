package classfile

type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableEntry
}

type LocalVariableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (c *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	c.LocalVariableTable = make([]*LocalVariableEntry, length)
	for i := range c.LocalVariableTable {
		c.LocalVariableTable[i] = &LocalVariableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
