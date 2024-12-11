package rtda

import (
	"github.com/stretchr/testify/assert"
	"gojvm/common/rtda/heap"
	"testing"
)

func TestLocalVars(t *testing.T) {
	tests := []struct {
		index uint
		typ   string
		val   any
	}{
		{0, "int", int32(100)},
		{1, "int", int32(-100)},
		{2, "long", int64(2997924580)},
		{4, "long", int64(-2997924580)},
		{6, "float", float32(3.1415926)},
		{7, "double", 2.7182818654645},
		{9, "ref", &heap.Object{}},
	}
	lvs := newLocalVars(10)
	for _, test := range tests {
		setLocalVarsValue(lvs, test.index, test.val)
	}
	for _, test := range tests {
		checkLocalVarsValue(t, lvs, test.index, test.typ, test.val)
	}
}

func setLocalVarsValue(lvs LocalVars, index uint, val any) {
	switch val.(type) {
	case int32:
		lvs.SetInt(index, val.(int32))
	case int64:
		lvs.SetLong(index, val.(int64))
	case float32:
		lvs.SetFloat(index, val.(float32))
	case float64:
		lvs.SetDouble(index, val.(float64))
	case *heap.Object:
		lvs.SetRef(index, val.(*heap.Object))
	default:
		panic("Illegal value")
	}
}

func checkLocalVarsValue(t *testing.T, lvs LocalVars, index uint, typ string, val any) {
	switch typ {
	case "int":
		assert.Equal(t, val, lvs.GetInt(index))
	case "long":
		assert.Equal(t, val, lvs.GetLong(index))
	case "float":
		assert.Equal(t, val, lvs.GetFloat(index))
	case "double":
		assert.Equal(t, val, lvs.GetDouble(index))
	case "ref":
		assert.Equal(t, val, lvs.GetRef(index))
	default:
		assert.FailNow(t, "Illegal type")
	}
}
