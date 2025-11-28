package main

import (
	"fmt"
	"unsafe"
)

type status = int8

const (
	white status = iota
	gray
	black
)

func Trace(stacks [][]uintptr) []uintptr {
	seen := make(map[uintptr]status, len(stacks)*len(stacks[0]))
	result := make([]uintptr, 0, len(stacks)*len(stacks[0]))
	// need to implement
	//dfs(stacks[0][4], &seen, &result)

	for _, stack := range stacks {
		for _, frame := range stack {
			dfs(frame, &seen, &result)
		}
	}

	return result
}

func dfs(pointer uintptr, seen *map[uintptr]status, result *[]uintptr) {
	fmt.Println(pointer)

	if pointer == 0 {
		return
	}

	if _, ok := (*seen)[pointer]; ok {
		return
	}

	(*result) = append((*result), pointer)
	(*seen)[pointer] = gray

	res := *(**int)(unsafe.Pointer(pointer))
	dfs(uintptr(unsafe.Pointer(res)), seen, result)

	(*seen)[pointer] = black
}
