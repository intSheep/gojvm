package heap

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (o *Object) Fields() Slots {
	return o.fields
}

// IsInstanceOf 判断对象是否是某个类的实例
// 判断对象是不是某个类的实例本质上是判断类能不能赋值给某个对象
func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
