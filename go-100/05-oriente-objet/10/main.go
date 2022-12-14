package main

import (
	"fmt"
)

type LinkNode[T any] struct {
	Value T
	Next  *LinkNode[T]
}

func NewLinkNode[T any](value T) *LinkNode[T] {
	return &LinkNode[T]{
		Value: value,
		Next:  nil,
	}
}

func main() {
	ln := NewLinkNode[int](1)
	fmt.Printf("Node %+v", ln)
}
