package tree

/*
leetcode 101. 对称二叉树
链接：https://leetcode-cn.com/problems/symmetric-tree

给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
进阶：
你可以运用递归和迭代两种方法解决这个问题吗？
*/

func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return compareTree(t1.Left, t2.Right) && compareTree(t1.Right, t2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compareTree(root.Left, root.Right)
}

// 迭代解法
func isSymmetric1(root *TreeNode) bool {
	queue := [][2]*TreeNode{{root, root}}

	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]

		if t[0] == nil && t[1] == nil {
			continue
		}

		if t[0] == nil || t[1] == nil {
			return false
		}

		if t[0].Val != t[1].Val {
			return false
		}

		queue = append(queue, [2]*TreeNode{t[0].Left, t[1].Right})
		queue = append(queue, [2]*TreeNode{t[0].Right, t[1].Left})
	}

	return true
}
