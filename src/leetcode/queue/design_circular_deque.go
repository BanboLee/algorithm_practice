package queue

/*
leetcode 641. 设计循环双端队列
链接：https://leetcode-cn.com/problems/design-circular-deque

设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。

示例：
MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4

提示：
所有值的范围为 [1, 1000]
操作次数的范围为 [1, 1000]
请不要使用内置的双端队列库。
*/

type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
}

type MyCircularDeque struct {
	head  *ListNode
	tail  *ListNode
	max   int
	count int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor1(k int) MyCircularDeque {
	return MyCircularDeque{nil, nil, k, 0}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.count == this.max {
		return false
	}
	if this.count == 0 {
		this.head = &ListNode{value, nil, nil}
		this.tail = this.head
	} else {
		this.head = &ListNode{value, this.head, nil}
		this.head.next.prev = this.head
	}
	this.count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.count == this.max {
		return false
	}

	if this.count == 0 {
		this.head = &ListNode{value, nil, nil}
		this.tail = this.head
	} else {
		this.tail.next = &ListNode{value, nil, this.tail}
		this.tail = this.tail.next
	}

	this.count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.count == 0 {
		return false
	}

	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.next
		this.head.prev = nil
	}
	this.count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.count == 0 {
		return false
	}

	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.tail = this.tail.prev
		this.tail.next = nil
	}
	this.count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.count == 0 {
		return -1
	}
	return this.head.val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.count == 0 {
		return -1
	}
	return this.tail.val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.max
}
