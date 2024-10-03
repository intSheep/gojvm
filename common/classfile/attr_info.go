package classfile

const (
	ATTRIBUTE_Code               = "Code"
	ATTRIBUTE_ConstantValue      = "ConstantValue"
	ATTRIBUTE_Deprecated         = "Deprecated"
	ATTRIBUTE_Exceptions         = "Exceptions"
	ATTRIBUTE_LineNumberTable    = "LineNumberTable"
	ATTRIBUTE_LocalVariableTable = "LocalVariableTableAttribute"
	ATTRIBUTE_SourceFile         = "SourceFile"
	ATTRIBUTE_Synthetic          = "Synthetic"
)

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	count := reader.readUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.readUint16()
	attrName := cp.getUtf8(nameIndex)
	length := reader.readUint32()
	attrInfo := newAttribute(attrName, length, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttribute(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case ATTRIBUTE_Code:
		return &CodeAttribute{cp: cp}
	case ATTRIBUTE_ConstantValue:
		return &ConstantValueAttribute{}
	case ATTRIBUTE_Deprecated:
		return &DeprecatedAttribute{}
	case ATTRIBUTE_Exceptions:
		return &ExceptionsAttribute{}
	case ATTRIBUTE_LineNumberTable:
		return &LineNumberTableAttribute{}
	case ATTRIBUTE_LocalVariableTable:
		return &LocalVariableTableAttribute{}
	case ATTRIBUTE_SourceFile:
		return &SourceFileAttribute{cp: cp}
	case ATTRIBUTE_Synthetic:
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
