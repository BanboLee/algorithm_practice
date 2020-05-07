package sliding_window

import (
	"math"
)

type ele struct {
	val  int
	next *ele
	prev *ele
}

type dequeue struct {
	head   *ele
	tail   *ele
	length int
}

func (dq *dequeue) pushBack(v int) *ele {
	e := &ele{v, nil, dq.tail}
	if dq.empty() {
		dq.head = e
	} else {
		dq.tail.next = e
		e.prev = dq.tail
	}
	dq.tail = e
	dq.length++
	return e
}

func (dq *dequeue) pushFront(v int) *ele {
	e := &ele{v, dq.head, nil}
	if dq.empty() {
		dq.tail = e
	} else {
		dq.head.prev = e
		e.next = dq.head
	}
	dq.head = e
	dq.length++
	return e
}

func (dq *dequeue) popBack() int {
	if dq.empty() {
		return math.MinInt32
	}

	e := dq.tail
	if dq.length == 1 {
		dq.head = nil
		dq.tail = nil
	} else {
		e.prev.next = nil
		dq.tail = e.prev
	}
	dq.length--
	return e.val
}

func (dq *dequeue) popFront() int {
	if dq.empty() {
		return math.MinInt32
	}
	e := dq.head
	if dq.length != 1 {
		e.next.prev = nil
		dq.head = e.next
	} else {
		dq.head = nil
		dq.tail = nil
	}
	dq.length--
	return e.val
}

func (dq *dequeue) empty() bool {
	return dq.length == 0
}
