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

type LinkedList[T any] struct {
	Head   *LinkNode[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head:   nil,
		length: 0,
	}
}

// InsertAtHead ajoute un élément en tête de liste
func (l *LinkedList[T]) InsertAtHead(value T) {
	//TODO
}

// Get retourne un élément par son index ou retourne une erreur
func (l *LinkedList[T]) Get(index int) (T, error) {
	//TODO
}

func main() {
	link := NewLinkedList[int]()
	link.InsertAtHead(4)
	link.InsertAtHead(3)
	link.InsertAtHead(2)
	link.InsertAtHead(1)

	v, _ := link.Get(2)
	fmt.Println("Element at index 2 is ", v)
}
