package heap

// IsAssignableFrom 判断c是否可以赋值给t
func (c *Class) isAssignableFrom(t *Class) bool {
	s := c
	// 当s和t相同的时候，可以进行赋值
	if s == t {
		return true
	}

	if !s.IsInterface() {
		// 如果s不是接口，那么t必须是s的子类
		return s.isSubClassOf(t)
	} else {
		// 如果s是接口，那么t必须实现s接口
		return s.isImplements(t)
	}
}

// isImplements 判断c是否实现了t接口
func (c *Class) isImplements(t *Class) bool {
	// 要判断c是否实现t的接口，需要看c类以及c的超类是否实现了t的接口
	// 具体而言需要从两个维度看 2 * 2 = 4
	// 1. c类是否实现了t的接口
	// 2. c类的超类是否实现了t的接口
	// 3. c类是否实现了t的子接口
	// 4. c类的超类是否实现了t的子接口
	for sc := c; sc != nil; sc = c.superClass {
		for _, i := range sc.interfaces {
			if i == t || i.isSubInterfaceOf(t) {
				return true
			}
		}
	}
	return false
}

// isSubInterfaceOf 判断c是否是t的子接口
func (c *Class) isSubInterfaceOf(t *Class) bool {
	for _, i := range c.interfaces {
		if i == t || i.isSubInterfaceOf(t) {
			return true
		}
	}
	return false
}
