package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCowBufferMine(t *testing.T) {
	str := "hello world"
	data := []byte(str)

	buffer := NewCOWBuffer(data)
	require.Equal(t, str, buffer.String())

	buffer1 := buffer.Clone()
	require.Equal(t, str, buffer1.String())
	require.Equal(t, 2, *buffer.counter)
	require.Equal(t, 2, *buffer1.counter)

	buffer2 := buffer.Clone()
	require.Equal(t, str, buffer2.String())
	require.Equal(t, 3, *buffer.counter)
	require.Equal(t, 3, *buffer1.counter)
	require.Equal(t, 3, *buffer2.counter)

	buffer2.Close()
	require.False(t, buffer2.Update(1, byte('E')))
	require.Equal(t, "", buffer2.String())
	require.Equal(t, 0, *buffer2.counter)
	buffer2.Close()
	require.False(t, buffer2.Update(1, byte('E')))

	require.Equal(t, str, buffer.String())
	require.Equal(t, str, buffer1.String())
	require.Equal(t, 2, *buffer.counter)
	require.Equal(t, 2, *buffer1.counter)

	require.True(t, buffer.Update(2, byte('T')))
	require.False(t, buffer.Update(200, byte('T')))
	require.False(t, buffer.Update(-200, byte('T')))
	require.Equal(t, 1, *buffer.counter)
	require.Equal(t, "heTlo world", buffer.String())

	require.Equal(t, 1, *buffer1.counter)
	require.Equal(t, str, buffer1.String())
}

func TestCOWBuffer(t *testing.T) {
	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1 := buffer.Clone()
	copy2 := buffer.Clone()

	assert.Equal(t, unsafe.SliceData(data), unsafe.SliceData(buffer.data))
	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	assert.True(t, (*byte)(unsafe.SliceData(data)) == unsafe.StringData(buffer.String()))
	assert.True(t, (*byte)(unsafe.StringData(buffer.String())) == unsafe.StringData(copy1.String()))
	assert.True(t, (*byte)(unsafe.StringData(copy1.String())) == unsafe.StringData(copy2.String()))

	assert.True(t, buffer.Update(0, 'g'))
	assert.False(t, buffer.Update(-1, 'g'))
	assert.False(t, buffer.Update(4, 'g'))

	assert.True(t, reflect.DeepEqual([]byte{'g', 'b', 'c', 'd'}, buffer.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy1.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy2.data))

	assert.NotEqual(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	copy1.Close()

	previous := copy2.data
	copy2.Update(0, 'f')
	current := copy2.data

	// 1 reference - don't need to copy buffer during update
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))

	copy2.Close()
}
