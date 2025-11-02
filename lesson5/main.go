package main

type possibleTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Map[T possibleTypes](data []T, action func(T) T) []T {
	if len(data) == 0 {
		return data
	}

	result := make([]T, len(data))

	for i := range data {
		result[i] = action(data[i])
	}

	return result
}

func Filter[T possibleTypes](data []T, action func(T) bool) []T {
	if len(data) == 0 {
		return data
	}

	result := make([]T, 0, len(data))

	for i := range data {
		if action(data[i]) {
			result = append(result, data[i])
		}
	}

	return result[:len(result):len(result)]
}

func Reduce[T possibleTypes](data []T, initial T, action func(T, T) T) T {
	result := initial

	for i := range data {
		result = action(result, data[i])
	}

	return result
}
