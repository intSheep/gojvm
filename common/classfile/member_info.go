package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i, _ := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}

func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}

// CodeAttribute 是类文件格式中的一个结构，它是method_info结构的一个可选属性，用于存储Java方法编译后的字节码指令和相关辅助信息。
// 当`.class`文件中的方法被编译成字节码时，与该方法相关的所有操作码、局部变量表、异常处理表等信息都会被封装在CodeAttribute中。
// 也就是说要执行一个方法就会读取要执行这个方法所需的操作码等信息
func (m *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range m.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (m *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range m.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
