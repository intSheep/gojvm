package heap

// SymRef 符号引用,存放运行时常量指针
type SymRef struct {
	cp        *ConstantPool
	className string // full type name
	class     *Class
}

func (s *SymRef) ResolvedClass() *Class {
	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

func (s *SymRef) resolveClassRef() {
	d := s.cp.class
	c := d.loader.LoadClass(s.className)
	if !c.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	s.class = c
}
