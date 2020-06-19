package tree

import "math"

/*
leetcode 124. 二叉树中的最大路径和
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum

给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:
输入: [1,2,3]
       1
      / \
     2   3
输出: 6
示例 2:
输入: [-10,9,20,null,null,15,7]
   -10
   / \
  9  20
    /  \
   15   7
输出: 42
*/

func helper124(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	lMax, rMax := helper124(root.Left, maxSum), helper124(root.Right, maxSum)

	nMax := max(max(root.Val+lMax, root.Val+rMax), max(root.Val+lMax, root.Val+lMax+rMax))
	if *maxSum < nMax {
		*maxSum = nMax
	}

	if *maxSum < root.Val {
		*maxSum = root.Val
	}

	return max(max(root.Val+rMax, root.Val+lMax), root.Val)
}

func maxPathSum(root *TreeNode) int {
	var res = math.MinInt64
	helper124(root, &res)

	return res
}
