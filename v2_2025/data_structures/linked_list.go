package main

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(value int) *LinkedList {
	newNode := NewNode(value)

	return &LinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value, next: nil}
}

func (node *Node) setNext(next *Node) {
	node.next = next
}

func (node *Node) setValue(value int) *Node {
	node.value = value
	return node
}

func main() {
}
