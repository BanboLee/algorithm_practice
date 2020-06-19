package tree

/*
leetcode 145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

func helper145(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = helper145(root.Left, res)
	res = helper145(root.Right, res)
	res = append(res, root.Val)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	return helper145(root, res)
}

// 迭代
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		res = append([]int{cur.Val}, res...)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return res
}
