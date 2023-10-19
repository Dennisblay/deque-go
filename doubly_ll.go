package main

import "fmt"

// Node represents a node in the doubly linked list
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// InsertAtBeginning inserts a new node with the given data at the beginning of the doubly linked list
func (dll *DoublyLinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{data: data}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		dll.size++
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
		dll.size++
	}
}

// InsertAtEnd inserts a new node with the given data at the end of the doubly linked list
func (dll *DoublyLinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{data: data}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		dll.size++
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
		dll.size++
	}
}

func (dll *DoublyLinkedList) PopFront() interface{} {
	if dll.IsEmpty() {
		return nil
	}
	nodeToRemove := dll.head
	if nodeToRemove == dll.tail {
		dll.head = nil
		dll.tail = nil
		dll.size--
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
		dll.size--
	}
	nodeToRemove.prev = nil
	nodeToRemove.next = nil
	return nodeToRemove.data

}

func (d *Deque) Front() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.head.data
}

func (dll *DoublyLinkedList) PopBack() interface{} {
	if dll.IsEmpty() {
		return nil
	}
	nodeToRemove := dll.tail
	if nodeToRemove == dll.head {
		dll.head = nil
		dll.tail = nil
		dll.size--
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		dll.size--
	}
	nodeToRemove.next = nil
	nodeToRemove.prev = nil
	return nodeToRemove.data
}

func (d *Deque) Back() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.tail.data
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

// DisplayForward prints the elements of the doubly linked list in forward order
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
}

// DisplayBackward prints the elements of the doubly linked list in backward order
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.prev
	}
	fmt.Println()
}

