package tree

/*
leetcode 116. 填充每个节点的下一个右侧节点指针
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node

给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。

提示：
你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 队列解法 O(N) + O(N)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left, cur.Right)
			}
			if i != l-1 {
				cur.Next = queue[i+1]
			}
		}
		queue = queue[l:]
	}

	return root
}

/*
   递归解法：O(N) + O(1)
        https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/java-san-xing-he-xin-dai-ma-chao-jian-ji-yi-yu-li-/
        每个 node 左子树的 next , 就是 node 的右子树
        每个 node 右子树的 next, 就是 node next 的 左子树
*/
func connect1(root *Node) *Node {
	dfs(root, nil)
	return root
}
func dfs(root, next *Node) {
	if root != nil {
		root.Next = next
		dfs(root.Left, root.Right)
		if next == nil {
			dfs(root.Right, nil)
		} else {
			dfs(root.Right, next.Left)
		}
	}
}

/*
   非递归O(N) + O(1)
*/
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	level := root
	for level != nil {
		head := level

		for head != nil {
			if head.Left != nil {
				head.Left.Next = head.Right
			}

			if head.Next != nil && head.Right != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}
		level = level.Left
	}

	return root
}
