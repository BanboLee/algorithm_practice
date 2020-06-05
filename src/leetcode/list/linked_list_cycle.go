package list

/*
leetcode 141. 环形链表
链接：https://leetcode-cn.com/problems/linked-list-cycle

给定一个链表，判断链表中是否有环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

进阶：
你能用 O(1)（即，常量）内存解决此问题吗？
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast, low := head, head
	for {
		if fast == nil {
			return false
		}
		if fast == low {
			return true
		}
	}
}

/*
leetcode 142. 环形链表 II
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。

进阶：
你是否可以不用额外空间解决此题？
*/

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		// 无环
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		// 此时 慢指针步数 Ss = m + n + k1l, 快指针步数为 Sf = m + n + k2l
		// 由此可得 m + n = kl, 当前慢指针已从入口处前进了n， 再前进m，必然回到入口处
		if fast == slow {
			break
		}
	}

	// 第一次相遇时必然是入口
	for slow != head {
		head = head.Next
		slow = slow.Next
	}
	return head
}
