package rtda

import "gojvm/common/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
