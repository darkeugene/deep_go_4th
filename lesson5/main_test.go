package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapMine(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expectedData := []int{1, 4, 9, 16, 25}

	require.Equal(t, expectedData, Map(data, func(datum int) int {
		return datum * datum
	}))

	expectedData = []int{-1, -2, -3, -4, -5}
	require.Equal(t, expectedData, Map(data, func(datum int) int {
		return -datum
	}))
}

func TestMapMineInt64(t *testing.T) {
	data := []int64{1, 2, 3, 4, 5}
	expectedData := []int64{1, 4, 9, 16, 25}

	require.Equal(t, expectedData, Map(data, func(datum int64) int64 {
		return datum * datum
	}))

	expectedData = []int64{-1, -2, -3, -4, -5}
	require.Equal(t, expectedData, Map(data, func(datum int64) int64 {
		return -datum
	}))
}

func TestFilterMine(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expectedData := []int{2, 4}

	require.Equal(t, expectedData, Filter(data, func(datum int) bool {
		return datum&1 == 0
	}))

	expectedData = []int{1, 3, 5}

	require.Equal(t, expectedData, Filter(data, func(datum int) bool {
		return datum&1 == 1
	}))
}

func TestFilterMineInt32(t *testing.T) {
	data := []int32{1, 2, 3, 4, 5}
	expectedData := []int32{2, 4}

	require.Equal(t, expectedData, Filter(data, func(datum int32) bool {
		return datum&1 == 0
	}))

	expectedData = []int32{1, 3, 5}

	require.Equal(t, expectedData, Filter(data, func(datum int32) bool {
		return datum&1 == 1
	}))
}

func TestReduceMine(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	require.Equal(t, 16, Reduce(data, 1, func(dat1 int, dat2 int) int {
		return dat1 + dat2
	}))

	require.Equal(t, 0, Reduce(data, 0, func(dat1 int, dat2 int) int {
		return dat1 * dat2
	}))
}

func TestReduceMineInt16(t *testing.T) {
	data := []int16{1, 2, 3, 4, 5}

	require.Equal(t, int16(16), Reduce(data, 1, func(dat1 int16, dat2 int16) int16 {
		return dat1 + dat2
	}))

	require.Equal(t, int16(0), Reduce(data, 0, func(dat1 int16, dat2 int16) int16 {
		return dat1 * dat2
	}))
}

func TestMap(t *testing.T) {
	tests := map[string]struct {
		data   []int
		action func(int) int
		result []int
	}{
		"nil numbers": {
			action: func(number int) int {
				return -number
			},
		},
		"empty numbers": {
			data: []int{},
			action: func(number int) int {
				return -number
			},
			result: []int{},
		},
		"inc numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(number int) int {
				return number + 1
			},
			result: []int{2, 3, 4, 5, 6},
		},
		"double numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(number int) int {
				return number * number
			},
			result: []int{1, 4, 9, 16, 25},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Map(test.data, test.action)
			assert.True(t, reflect.DeepEqual(test.result, result))
		})
	}
}

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		data   []int
		action func(int) bool
		result []int
	}{
		"nil numbers": {
			action: func(number int) bool {
				return number == 0
			},
		},
		"empty numbers": {
			data: []int{},
			action: func(number int) bool {
				return number == 1
			},
			result: []int{},
		},
		"even numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(number int) bool {
				return number%2 == 0
			},
			result: []int{2, 4},
		},
		"positive numbers": {
			data: []int{-1, -2, 1, 2},
			action: func(number int) bool {
				return number > 0
			},
			result: []int{1, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Filter(test.data, test.action)
			assert.True(t, reflect.DeepEqual(test.result, result))
		})
	}
}

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		initial int
		data    []int
		action  func(int, int) int
		result  int
	}{
		"nil numbers": {
			action: func(lhs, rhs int) int {
				return 0
			},
		},
		"empty numbers": {
			data: []int{},
			action: func(lhs, rhs int) int {
				return 0
			},
		},
		"sum of numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(lhs, rhs int) int {
				return lhs + rhs
			},
			result: 15,
		},
		"sum of numbers with initial value": {
			initial: 10,
			data:    []int{1, 2, 3, 4, 5},
			action: func(lhs, rhs int) int {
				return lhs + rhs
			},
			result: 25,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Reduce(test.data, test.initial, test.action)
			assert.Equal(t, test.result, result)
		})
	}
}
