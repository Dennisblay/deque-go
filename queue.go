package main

import (
	"fmt"
)

type Queue struct {
	LinkedList
}

func (q *Queue) enqueue(data interface{}) {
	q.AddNode(data)
}

func (q *Queue) dequeue() interface{} {
	front := q.front()
	if !q.isEmpty() {
		q.Remove(q.head.data)
	}
	return front

}

func (q *Queue) front() interface{} {
	if q.isEmpty() {
		return nil
	}
	return q.head.data
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}
