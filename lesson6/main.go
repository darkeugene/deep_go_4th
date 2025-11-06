package main

import (
	"fmt"
	"unsafe"
)

func main() {
	g := GamePerson{}
	fmt.Println("Size of struct: ", unsafe.Sizeof(g))
	fmt.Println("Align: ", unsafe.Alignof(g))
}
