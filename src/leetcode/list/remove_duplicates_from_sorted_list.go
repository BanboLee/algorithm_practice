package list

/*
leetcode 83. 删除排序链表中的重复元素
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list

给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:
输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur.Next != nil {
		if cur.Next.Val != cur.Val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}
	}

	return head
}

/*
leetcode 82. 删除排序链表中的重复元素 II
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:
输入: 1->1->1->2->3
输出: 2->3
*/

//	强行按照一般思想硬撸
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := &ListNode{head.Val - 1, nil}
	nh := prev
	cur := head
	for cur != nil {
		n := cur.Next
		if n == nil || cur.Val != n.Val {
			prev.Next = cur
			prev = prev.Next
			cur = n
			continue
		}

		for n != nil && cur.Val == n.Val {
			n = n.Next
		}

		cur = n
	}
	prev.Next = nil
	return nh.Next
}
