package main

import "fmt"

type LinkedList[T any] struct {
	head, tail *Node[T]
	len        int
}

type Node[T any] struct {
	next  *Node[T]
	value T
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

/*
	to define method on generic type we can use the same syntax but consider that LinkedList type
	has an extension of type parameter -> LinkedList[T] that's the type now
*/

func (ll *LinkedList[T]) Push(value T) {
	if ll.tail == nil {
		ll.head = NewNode(value)
		ll.tail = ll.head
	} else {
		ll.tail.next = NewNode(value)
		ll.tail = ll.tail.next
	}
}

func (ll *LinkedList[T]) AllElements() []T {
	var elems []T
	for e := ll.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}

	return elems
}

func main() {
	ll := LinkedList[float64]{}

	for i := 0.0; i < 5.0; i += 0.1 {
		ll.Push(i)
	}

	fmt.Println("LinkedList", ll.AllElements())
}
