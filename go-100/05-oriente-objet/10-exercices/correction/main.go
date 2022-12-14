package main

import (
	"errors"
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
	newNode := NewLinkNode(value)
	newNode.Next = l.Head
	l.Head = newNode
	l.length++
}

// Get retourne un élément par son index ou retourne une erreur
func (l *LinkedList[T]) Get(index int) (T, error) {
	// gestion du out of range
	var t T
	if index < 0 || index > l.length-1 {
		return t, errors.New("index out of bound")
	}
	// cas general: trouver l'élément au bon index et retourner sa valeur
	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Value, nil
}

func main() {
	link := NewLinkedList[int]()
	link.InsertAtHead(4)
	link.InsertAtHead(3)
	link.InsertAtHead(2)
	link.InsertAtHead(1)

	link2 := NewLinkedList[string]()
	link2.InsertAtHead("d")
	link2.InsertAtHead("c")
	link2.InsertAtHead("b")
	link2.InsertAtHead("a")

	v, _ := link.Get(2)
	fmt.Println("Element at index 2 is ", v)

	v2, _ := link2.Get(2)
	fmt.Println("Element at index 2 is ", v2)
}
