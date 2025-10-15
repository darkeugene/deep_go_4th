package main

import (
	"fmt"
	"unsafe"
)

func ToLittleEndian[T ~uint64 | ~uint32 | ~uint16 | ~uint8](number T) T {
	size := int(unsafe.Sizeof(number))
	pointer := unsafe.Pointer(&number)

	for i := 0; i < size/2; i++ {
		firstByte := (*int8)(unsafe.Add(pointer, i))
		lastByte := (*int8)(unsafe.Add(pointer, size-i-1))

		*firstByte, *lastByte = *lastByte, *firstByte
	}

	return number
}

func main() {
	var num1 uint32 = 0x03040506

	number := ToLittleEndian(num1)

	fmt.Printf("%x", num1)
	fmt.Println()
	fmt.Printf("%x", number)
}
