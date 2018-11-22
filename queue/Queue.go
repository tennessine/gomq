package queue

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.head.value, true
}

func (q *Queue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (q *Queue) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *Queue) Remove() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value, true
}
