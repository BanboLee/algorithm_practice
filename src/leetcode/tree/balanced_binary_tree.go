package tree

/*
leetcode 110. 平衡二叉树
链接：https://leetcode-cn.com/problems/balanced-binary-tree

给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func balanceHelper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	lb, ll := balanceHelper(root.Left)
	rb, rl := balanceHelper(root.Right)

	if !lb || !rb {
		return false, 0
	}

	if abs(ll-rl) > 1 {
		return false, 0
	}
	return true, max(ll, rl) + 1
}

func isBalanced(root *TreeNode) bool {
	b, _ := balanceHelper(root)
	return b
}
