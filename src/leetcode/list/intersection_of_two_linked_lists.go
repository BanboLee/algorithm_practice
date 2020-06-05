package list

/*
leetcode 160. 相交链表
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists

编写一个程序，找到两个单链表相交的起始节点。

注意：
如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
*/

/*
   1. 可以计算长度，然后让长的先过去一部分
   2. 让两个指针走过相同的路程，速度一样，最终都会到达同一个终点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next

		}
	}
	return p1
}
