package list

/*
leetcode 21.合并两个有序的链表
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	} else {
		head = l1
		l1 = l1.Next
	}
	cur := head
	for {
		if l1 == nil {
			cur.Next = l2
			return head
		}
		if l2 == nil {
			cur.Next = l1
			return head
		}
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
}

// 递归解法
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	// base case
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 拆分问题	合并结果
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}
