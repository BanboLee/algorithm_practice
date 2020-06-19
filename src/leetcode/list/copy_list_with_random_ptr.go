package list

/*
leetcode 138. 复制带随机指针的链表
链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer

给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
要求返回这个链表的 深拷贝。
我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
示例 4：
输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。

提示：
-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 回溯法
var mp = make(map[*Node]*Node) // map 中使得原节点和新节点一一对应
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	if v, ok := mp[head]; ok {
		return v
	}

	// copy
	node := &Node{head.Val, nil, nil}
	mp[head] = node

	node.Next = copyRandomList(head.Next)
	node.Random = copyRandomList(head.Random)
	return node
}

// O(N)迭代法  不用递归
func copyNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := mp[node]; ok {
		return v
	}

	newNode := &Node{node.Val, nil, nil}
	mp[node] = newNode
	return newNode
}
func copyRandomList1(head *Node) *Node {
	var dummy = &Node{-1, nil, nil}
	cur := dummy
	for head != nil {
		cur.Next = copyNode(head)
		cur = cur.Next
		cur.Random = copyNode(head.Random)
		head = head.Next
	}

	return dummy.Next
}

// O(1)迭代法
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	// step1. get A` for A
	cur := head
	for cur != nil {
		cur.Next = &Node{cur.Val, cur.Next, nil}
		cur = cur.Next.Next
	}

	// step2. get random
	cur = head
	for cur != nil {
		newNode := cur.Next
		if cur.Random != nil {
			newNode.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// step3. get new list
	dummy := &Node{-1, nil, nil}
	cur = dummy
	for head != nil {
		cur.Next = head.Next
		cur = cur.Next
		head.Next = cur.Next
		head = head.Next
	}

	return dummy.Next
}
