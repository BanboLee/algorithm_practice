package tree

/*
leetcode 94. 二叉树的中序遍历
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal

给定一个二叉树，返回它的中序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 迭代, 一直按顺序压栈然后弹出
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	var res []int

	for root != nil || len(stack) != 0 {
		// 中序遍历要把所有左子树先全都压进去
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 当不再有左子树的时候，遍历自己
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 然后遍历右子树
		root = root.Right
	}
	return res
}
