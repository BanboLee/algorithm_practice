package list

/*
leetcode 61. 旋转链表
链接：https://leetcode-cn.com/problems/rotate-list

给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// step1. calc length
	l := 0
	tail := head
	for tail.Next != nil {
		l++
		tail = tail.Next
	}

	// step2. calc nums of steps.
	remain := k
	if l == k {
		return head
	} else if l < k {
		remain = l - (k % l)
	} else {
		remain = l - k
	}

	// step3. rotate
	tail.Next = head
	for ; remain > 0; remain-- {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil

	return head
}