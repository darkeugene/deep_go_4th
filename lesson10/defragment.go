package main

import (
	"unsafe"
)

func Defragment[T byte | uint16 | uint32 | uint64](memory []byte, pointers []unsafe.Pointer) {
	var example T
	blockSize := int(unsafe.Sizeof(example))

	if len(memory) <= len(pointers)*blockSize {
		return
	}

	pointersSet := make(map[uintptr]struct{}, len(pointers))
	for _, pointer := range pointers {
		pointersSet[uintptr(pointer)] = struct{}{}
	}

	currentPointer := unsafe.Pointer(&memory[0])
	currentPointerIndex := 0

	for {
		if len(pointersSet) == 0 {
			return
		}

		if _, ok := pointersSet[uintptr(pointers[currentPointerIndex])]; !ok {
			currentPointerIndex++
			continue
		}

		if _, ok := pointersSet[uintptr(currentPointer)]; ok {
			delete(pointersSet, uintptr(currentPointer))
			currentPointer = unsafe.Add(currentPointer, blockSize)

			continue
		}

		srcBytes := unsafe.Slice((*byte)(pointers[currentPointerIndex]), blockSize)
		dstBytes := unsafe.Slice((*byte)(currentPointer), blockSize)

		copy(dstBytes, srcBytes)
		clear(srcBytes)

		delete(pointersSet, uintptr(pointers[currentPointerIndex]))
		pointers[currentPointerIndex] = currentPointer
		currentPointer = unsafe.Add(currentPointer, blockSize)
		currentPointerIndex++
	}
}
