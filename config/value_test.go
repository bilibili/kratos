package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_atomicValue_Bool(t *testing.T) {
	var vlist []interface{}
	vlist = []interface{}{"1", "t", "T", "true", "TRUE", "True", true}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Bool()
		assert.NoError(t, err, b)
		assert.True(t, b, b)
	}

	vlist = []interface{}{"0", "f", "F", "false", "FALSE", "False", false}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Bool()
		assert.NoError(t, err, b)
		assert.False(t, b, b)
	}

	vlist = []interface{}{int32(1), 1, uint16(1), "bbb", "-1"}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Bool()
		assert.Error(t, err, b)
	}
}

func Test_atomicValue_Int(t *testing.T) {
	var vlist []interface{}
	vlist = []interface{}{"123123", float64(123123), int64(123123)}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Int()
		assert.NoError(t, err, b)
		assert.Equal(t, int64(123123), b, b)
	}

	vlist = []interface{}{int32(1123123), 123131, uint16(1), "bbb", "-x1"}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Int()
		assert.Error(t, err, b)
	}
}

func Test_atomicValue_Float(t *testing.T) {
	var vlist []interface{}
	vlist = []interface{}{"123123.1", float64(123123.1)}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Float()
		assert.NoError(t, err, b)
		assert.Equal(t, float64(123123.1), b, b)
	}

	vlist = []interface{}{float32(1123123), 123131, uint16(1), "bbb", "-x1"}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Float()
		assert.Error(t, err, b)
	}
}

func Test_atomicValue_String(t *testing.T) {
	var vlist []interface{}
	vlist = []interface{}{"1", float64(1), int64(1)}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.String()
		assert.NoError(t, err, b)
		assert.Equal(t, "1", b, b)
	}

	v := atomicValue{}
	v.Store(true)
	b, err := v.String()
	assert.NoError(t, err, b)
	assert.Equal(t, "true", b, b)
}

func Test_atomicValue_Duration(t *testing.T) {
	var vlist []interface{}
	vlist = []interface{}{5*time.Second}
	for _, x := range vlist {
		v := atomicValue{}
		v.Store(x)
		b, err := v.Duration()
		assert.NoError(t, err)
		assert.Equal(t, 5*time.Second, b)
	}
}


func Test_atomicValue_Scan(t *testing.T) {
	var err error
	v := atomicValue{}
	err = v.Scan(&struct {
		A string `json:"a"`
	}{"a"})
	assert.NoError(t, err)

	err = v.Scan(&struct {
		A string `json:"a"`
	}{"a"})
	assert.NoError(t, err)
}


