package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.queue))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.queue))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}

func TestCircularQueue3(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.False(t, queue.Pop())
	assert.True(t, queue.Push(1))
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 1, queue.Back())

	assert.True(t, queue.Push(4))
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Push(7))
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 7, queue.Back())

	assert.False(t, queue.Push(9))
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 7, queue.Back())

	assert.True(t, queue.Pop())
	assert.Equal(t, 4, queue.Front())
	assert.Equal(t, 7, queue.Back())

	assert.True(t, queue.Push(22))
	assert.Equal(t, 4, queue.Front())
	assert.Equal(t, 22, queue.Back())

	assert.True(t, queue.Pop())
	assert.Equal(t, 7, queue.Front())
	assert.Equal(t, 22, queue.Back())

	assert.True(t, queue.Pop())
	assert.Equal(t, 22, queue.Front())
	assert.Equal(t, 22, queue.Back())

	assert.True(t, queue.Pop())
	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())

	assert.False(t, queue.Pop())
	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
}

func TestCircularQueue0(t *testing.T) {
	const queueSize = 0
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.False(t, queue.Pop())
	assert.False(t, queue.Push(1))
	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
}

func TestCircularQueue08(t *testing.T) {
	const queueSize = 0
	queue := NewCircularQueue[int8](queueSize)

	assert.True(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.False(t, queue.Pop())
	assert.False(t, queue.Push(1))
	assert.Equal(t, int8(-1), queue.Front())
	assert.Equal(t, int8(-1), queue.Back())
}

func TestCircularQueue016(t *testing.T) {
	const queueSize = 0
	queue := NewCircularQueue[int16](queueSize)

	assert.True(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.False(t, queue.Pop())
	assert.False(t, queue.Push(1))
	assert.Equal(t, int16(-1), queue.Front())
	assert.Equal(t, int16(-1), queue.Back())
}

func TestCircularQueue032(t *testing.T) {
	const queueSize = 0
	queue := NewCircularQueue[int32](queueSize)

	assert.True(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.False(t, queue.Pop())
	assert.False(t, queue.Push(1))
	assert.Equal(t, int32(-1), queue.Front())
	assert.Equal(t, int32(-1), queue.Back())
}

func TestCircularQueue064(t *testing.T) {
	const queueSize = 0
	queue := NewCircularQueue[int64](queueSize)

	assert.True(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.False(t, queue.Pop())
	assert.False(t, queue.Push(1))
	assert.Equal(t, int64(-1), queue.Front())
	assert.Equal(t, int64(-1), queue.Back())
}
