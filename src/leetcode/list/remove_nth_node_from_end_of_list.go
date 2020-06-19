package list

/*
leetcode 19. 删除链表的倒数第N个节点
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.

说明：
给定的 n 保证是有效的。

进阶：
你能尝试使用一趟扫描实现吗？
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var mp []*ListNode
	cnt := 0
	cur := head
	for cur != nil {
		mp = append(mp, cur)
		cur = cur.Next
		cnt++
	}

	if cnt <= 1 {
		return nil
	}

	idx := cnt - n
	if idx == 0 {
		return head.Next
	}
	mp[idx-1].Next = mp[idx].Next
	return head
}

// 双指针法, 慢指针总是比快指针慢 n - 1 步
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	p1, p2, dis := head, head, 0
	for p2 != nil {
		p2 = p2.Next
		if dis >= n+1 {
			p1 = p1.Next
		}
		dis++
	}

	if dis == n {
		head = p1.Next
	} else {
		p1.Next = p1.Next.Next
	}
	return head
}

// 哑节点 + 双指针
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	p1, p2 := dummy, dummy
	for i := 0; i <= n; i++ {
		p2 = p2.Next
	}

	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next
	return dummy.Next
}
