package tree

/*
leetcode 144. 二叉树的前序遍历
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal

给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res = []int{root.Val}

	if root.Left == nil && root.Right == nil {
		return res
	}

	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}

// 迭代实现
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		res = append(res, cur.Val)
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
	}

	return res
}
