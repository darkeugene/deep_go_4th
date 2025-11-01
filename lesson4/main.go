package main

import "fmt"

func main() {
	orderedMap := NewOrderedMap[int, int]()
	fmt.Println(orderedMap.Size())
}
