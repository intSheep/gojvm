package classfile

// 行号信息
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (c *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	c.lineNumberTable = make([]*LineNumberTableEntry, length)
	for i := range c.lineNumberTable {
		c.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
