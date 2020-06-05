package list

/*
leetcode 206. 翻转列表
链接：https://leetcode-cn.com/problems/reverse-linked-list

反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListFromSlice(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	var head = &ListNode{data[0], nil}
	var cur = head
	for i := 1; i < len(data); i++ {
		cur.Next = &ListNode{data[i], nil}
		cur = cur.Next
	}
	return head
}

func ListToSlice(head *ListNode) []int {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	return data
}

func reverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode
		cur  *ListNode
	)

	prev = nil
	cur = head

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

// 递归解法
func reverseList1(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}

	// 差分问题，进而解决head以后的链表
	tmp := reverseList(head.Next)

	// 结果组合
	// next 指向head
	head.Next.Next = head
	// head 此时变成尾部
	head.Next = nil
	return tmp
}
