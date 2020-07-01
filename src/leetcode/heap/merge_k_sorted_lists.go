package heap

/*
leetcode 23. 合并K个排序链表
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func CompareListNode(node1, node2 interface{}) bool {
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.(*ListNode).Val < node2.(*ListNode).Val
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pq := &PriorityQueue{nil, CompareListNode}
	for _, list := range lists {
		pq.Insert(list)
	}

	var dummy = &ListNode{-1, nil}
	head := dummy
	for pq.Pop() != nil {
		head.Next = pq.Pop().(*ListNode)
		head = head.Next
		if head.Next != nil {
			pq.Insert(head.Next)
		}
		head.Next = nil
	}
	return dummy.Next
}
