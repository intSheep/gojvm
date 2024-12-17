package heap

import "gojvm/common/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (m *MemberRef) copyMemberRef(refInfo *classfile.ConstantMemberRefInfo) {
	m.className = refInfo.ClassName()
	m.name, m.descriptor = refInfo.NameAndDescriptor()
}

func (m *MemberRef) Name() string {
	return m.name
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
}
