// package main
//
//	type Deque struct {
//		front *Node
//		back  *Node
//		Size  int // Size field to keep track of the deque size
//	}
//
//	func (dq *Deque) PushFront(point Point) {
//		newNode := &Node{point: point}
//
//		if dq.front == nil {
//			dq.front = newNode
//			dq.back = newNode
//		} else {
//			newNode.next = dq.front
//			dq.front.prev = newNode
//			dq.front = newNode
//		}
//
//		dq.Size++ // Increment the size after adding an element
//	}
//
//	func (dq *Deque) PopFront() (Point, bool) {
//		if dq.front == nil {
//			return Point{}, false
//		}
//
//		point := dq.front.point
//		dq.front = dq.front.next
//
//		if dq.front == nil {
//			dq.back = nil
//		} else {
//			dq.front.prev = nil
//		}
//
//		dq.Size-- // Decrement the size after removing an element
//		return point, true
//	}
//
//	func (dq *Deque) PushBack(point Point) {
//		newNode := &Node{point: point}
//
//		if dq.back == nil {
//			dq.front = newNode
//			dq.back = newNode
//		} else {
//			newNode.prev = dq.back
//			dq.back.next = newNode
//			dq.back = newNode
//		}
//
//		dq.Size++ // Increment the size after adding an element
//	}
//
//	func (dq *Deque) PopBack() (Point, bool) {
//		if dq.back == nil {
//			return Point{}, false
//		}
//
//		point := dq.back.point
//		dq.back = dq.back.prev
//
//		if dq.back == nil {
//			dq.front = nil
//		} else {
//			dq.back.next = nil
//		}
//
//		dq.Size-- // Decrement the size after removing an element
//		return point, true
//	}
//
//	func (dq *Deque) IsEmpty() bool {
//		return dq.front == nil
//	}
package main

import "fmt"

// Deque This a double ended queue implementation
type Deque struct {
	items []interface{}
}

func (dq *Deque) IsEmpty() bool {
	return len(dq.items) == 0
}

func (dq *Deque) Size() int {
	return len(dq.items)
}

func (dq *Deque) AddFront(item interface{}) {
	dq.items = append([]interface{}{item}, dq.items...)
}

func (dq *Deque) AddRear(item interface{}) {
	dq.items = append(dq.items, item)
}

func (dq *Deque) RemoveFront() interface{} {
	if dq.IsEmpty() {
		return nil
	}
	front := dq.items[0]
	dq.items = dq.items[1:]
	return front
}

func (dq *Deque) RemoveRear() interface{} {
	if dq.IsEmpty() {
		return nil
	}
	rear := dq.items[len(dq.items)-1]
	dq.items = dq.items[:len(dq.items)-1]
	return rear
}

func main_() {
	dq := Deque{}

	dq.AddFront(1)
	dq.AddFront(2)
	dq.AddRear(3)

	fmt.Println("Deque size:", dq.Size())

	front := dq.RemoveFront()
	rear := dq.RemoveRear()

	fmt.Println("Front:", front)
	fmt.Println("Rear:", rear)
	fmt.Println("Is Deque empty?", dq.IsEmpty())
}
