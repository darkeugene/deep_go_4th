package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap[int, int]()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	assert.Equal(t, 6, data.Size())
	data.ForEach(func(key, _ int) {})

	data.Erase(14)
	assert.Equal(t, 5, data.Size())
	data.ForEach(func(key, _ int) {})

	data.Erase(2)
	assert.Equal(t, 4, data.Size())
	data.ForEach(func(key, _ int) {})

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}

func TestOrderedMap(t *testing.T) {
	data := NewOrderedMap[int, int]()
	assert.Equal(t, 0, data.Size())

	data.Insert(33, 66)
	assert.Equal(t, 1, data.Size())

	assert.Equal(t, true, data.Contains(33))
	assert.Equal(t, false, data.Contains(66))

	data.Insert(5, 10)
	data.Insert(1, 2)
	data.Insert(4, 8)
	data.Insert(20, 40)
	data.Insert(17, 34)
	data.Insert(31, 62)

	assert.Equal(t, true, data.Contains(33))
	assert.Equal(t, true, data.Contains(1))
	assert.Equal(t, true, data.Contains(4))
	assert.Equal(t, true, data.Contains(20))
	assert.Equal(t, true, data.Contains(17))
	assert.Equal(t, true, data.Contains(31))

	assert.Equal(t, false, data.Contains(0))

	data.Insert(35, 70)
	data.Insert(99, 198)

	assert.Equal(t, true, data.Contains(35))
	assert.Equal(t, true, data.Contains(99))

	assert.Equal(t, 9, data.Size())

	keys := []int{}
	values := []int{}
	expectedKeys := []int{1, 4, 5, 17, 20, 31, 33, 35, 99}
	expectedValues := []int{2, 8, 10, 34, 40, 62, 66, 70, 198}
	data.ForEach(func(key, value int) {
		keys = append(keys, key)
		values = append(values, value)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
	assert.True(t, reflect.DeepEqual(expectedValues, values))

	data.Erase(35)

	assert.Equal(t, 8, data.Size())

	data.Erase(20)

	assert.Equal(t, 7, data.Size())

	data.Erase(5)

	assert.Equal(t, 6, data.Size())

	keys = []int{}
	values = []int{}
	expectedKeys = []int{1, 4, 17, 31, 33, 99}
	expectedValues = []int{2, 8, 34, 62, 66, 198}
	data.ForEach(func(key, value int) {
		keys = append(keys, key)
		values = append(values, value)
	})
	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
	assert.True(t, reflect.DeepEqual(expectedValues, values))
}

func TestOrderedMapStr(t *testing.T) {
	data := NewOrderedMap[string, any]()
	assert.Equal(t, 0, data.Size())

	data.Insert("hello", "world")
	assert.Equal(t, 1, data.Size())

	assert.Equal(t, true, data.Contains("hello"))
	assert.Equal(t, false, data.Contains("b"))

	data.Insert("a", 66)
	data.Insert("e", 2e5)
	data.Insert("b", "b")
	data.Insert("c", 4)
	data.Insert("bc", 3.12)
	data.Insert("ab", "str")

	assert.Equal(t, true, data.Contains("a"))
	assert.Equal(t, true, data.Contains("b"))
	assert.Equal(t, true, data.Contains("ab"))
	assert.Equal(t, true, data.Contains("bc"))
	assert.Equal(t, true, data.Contains("c"))
	assert.Equal(t, true, data.Contains("e"))
	assert.Equal(t, true, data.Contains("hello"))

	assert.Equal(t, false, data.Contains("f"))

	data.Insert("td", 0.001)
	data.Insert("dt", -8)

	assert.Equal(t, true, data.Contains("dt"))
	assert.Equal(t, true, data.Contains("td"))

	assert.Equal(t, 9, data.Size())

	keys := []string{}
	values := []any{}
	expectedKeys := []string{"a", "ab", "b", "bc", "c", "dt", "e", "hello", "td"}
	expectedValues := []any{66, "str", "b", 3.12, 4, -8, 2e5, "world", 0.001}
	data.ForEach(func(key string, value any) {
		keys = append(keys, key)
		values = append(values, value)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
	assert.True(t, reflect.DeepEqual(expectedValues, values))

	data.Erase("c")

	assert.Equal(t, 8, data.Size())

	data.Erase("e")

	assert.Equal(t, 7, data.Size())

	data.Erase("bc")

	assert.Equal(t, 6, data.Size())

	keys = []string{}
	values = []any{}
	expectedKeys = []string{"a", "ab", "b", "dt", "hello", "td"}
	expectedValues = []any{66, "str", "b", -8, "world", 0.001}
	data.ForEach(func(key string, value any) {
		keys = append(keys, key)
		values = append(values, value)
	})
	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
	assert.True(t, reflect.DeepEqual(expectedValues, values))
}
