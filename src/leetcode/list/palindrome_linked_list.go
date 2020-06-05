package list

/*
leecode 234. 回文链表
链接：https://leetcode-cn.com/problems/palindrome-linked-list

请判断一个链表是否为回文链表。

示例 1:
输入: 1->2
输出: false
示例 2:
输入: 1->2->2->1
输出: true

进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/*
暴力法
正向遍历得到一个slice
再遍历一次反向比较slice
*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var data []int
	cur := head
	for cur != nil {
		data = append(data, cur.Val)
		cur = cur.Next
	}

	cur = head
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != cur.Val {
			return false
		}
		cur = cur.Next
	}

	if cur != nil {
		return false
	}

	return true
}

/*
   快慢指针+翻转后半部分
   1. 快慢指针找到中点
   2. 翻转后半部分
   3. 对比求解
   4。恢复链表
*/
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找到中点
	n, p := head, head
	for p.Next != nil && p.Next.Next != nil {
		n = n.Next
		p = p.Next.Next
	}

	// 翻转后半部分
	var prev *ListNode
	cur := n.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	// 求解
	var res = true
	back := prev
	for back != nil {
		if back.Val != head.Val {
			res = false
			break
		}
		back = back.Next
		head = head.Next
	}

	// 恢复
	var pp *ListNode
	for prev != nil {
		tmp := prev.Next
		prev.Next = pp
		pp = prev
		prev = tmp
	}
	n.Next = pp

	return res
}
