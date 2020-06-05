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

// 溢出  雪崩
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1 int64
	var v2 int64
	var c int64 = 1
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 += c * int64(l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			v2 += c * int64(l2.Val)
			l2 = l2.Next
		}
		c *= 10
	}

	sum := v1 + v2
	println(v1, v2, sum)
	if sum == 0 {
		return &ListNode{0, nil}
	}
	p := &ListNode{-1, nil}
	c = 1
	cur := p
	for c <= sum {
		cur.Next = &ListNode{int(sum / c % 10), nil}
		cur = cur.Next
		c *= 10
	}
	return p.Next
}

// TODO
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
