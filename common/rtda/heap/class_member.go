package heap

import "gojvm/common/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string //字段描述符
	class       *Class
}

// memberInfo从内存到运行时数据
func (c *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	c.accessFlags = memberInfo.AccessFlags()
	c.name = memberInfo.Name()
	c.descriptor = memberInfo.Descriptor()
}

func (c *ClassMember) Descriptor() string {
	return c.descriptor
}

func (c *ClassMember) Name() string {
	return c.name
}

func (c *ClassMember) Class() *Class {
	return c.class
}
