package classfile

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error parsing class file: %v", r)
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, nil
}

func (cf *ClassFile) read(cr *ClassReader) {
	cf.readAndCheckMagic(cr)
	cf.readAndCheckVersion(cr)
	cf.readConstantPool(cr)
	cf.accessFlags = cr.readUint16()
	cf.thisClass = cr.readUint16()
	cf.superClass = cr.readUint16()
	cf.interfaces = cr.readUint16s()
	cf.fields = readMembers(cr, cf.constantPool)
	cf.methods = readMembers(cr, cf.constantPool)
	cf.attributes = readAttributes(cr, cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(cr *ClassReader) {
	magic := cr.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(cr *ClassReader) {
	cf.minorVersion = cr.readUint16()
	cf.majorVersion = cr.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	default:
		logrus.Infof("java version :%v , maybe not support", cf.majorVersion)
	}
}

func (cf *ClassFile) readConstantPool(cr *ClassReader) {
	cf.constantPool = readConstantPool(cr)
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}
func (cf *ClassFile) InterfaceNames() []string {
	interfaces := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaces[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaces
}

func (cf *ClassFile) Show(isShow ...bool) {
	if len(isShow) == 0 {
		fmt.Println(cf.String())
	} else {
		if isShow[0] {
			fmt.Println(cf.String())
		}
	}
}

func (cf *ClassFile) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion()))
	buffer.WriteString(fmt.Sprintf("Constant Count:%v\n", len(cf.ConstantPool())))
	buffer.WriteString(fmt.Sprintf("Access Flags: 0x%x\n", cf.AccessFlags()))
	buffer.WriteString(fmt.Sprintf("This Class: %v\n", cf.ClassName()))
	buffer.WriteString(fmt.Sprintf("Super Class: %v\n", cf.SuperClassName()))
	buffer.WriteString(fmt.Sprintf("Interfaces: %v\n", cf.InterfaceNames()))
	buffer.WriteString(fmt.Sprintf("Fields Count: %v\n", len(cf.Fields())))
	for _, f := range cf.Fields() {
		buffer.WriteString(fmt.Sprintf("   %s\n", f.Name()))
	}
	buffer.WriteString(fmt.Sprintf("Methods Count: %v\n", len(cf.Methods())))
	for _, m := range cf.Methods() {
		buffer.WriteString(fmt.Sprintf("   %s\n", m.Name()))
	}
	return buffer.String()
}
