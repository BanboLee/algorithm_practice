package list

/*
leetcode 2. 两数相加
链接：https://leetcode-cn.com/problems/add-two-numbers

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0 // carry
	head := &ListNode{-1, nil}
	prev := head
	for l1 != nil || l2 != nil {
		var v int
		if l1 != nil && l2 != nil {
			v = l1.Val + l2.Val + c
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil {
			v = l2.Val + c
			l2 = l2.Next
		} else {
			v = l1.Val + c
			l1 = l1.Next
		}

		c = v / 10
		prev.Next = &ListNode{v % 10, nil}
		prev = prev.Next
	}

	if c == 1 {
		prev.Next = &ListNode{1, nil}
	}

	return head.Next
}
