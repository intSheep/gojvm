package classfile

type DeprecatedAttribute struct {
	MakerAttribute
}

type SyntheticAttribute struct {
	MakerAttribute
}

type MakerAttribute struct{}

func (c *MakerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
