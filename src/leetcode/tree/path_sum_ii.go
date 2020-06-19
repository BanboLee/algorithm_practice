package tree

/*
leetcode 113. 路径总和 II
链接：https://leetcode-cn.com/problems/path-sum-ii

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/

func dfs1(root *TreeNode, sum int, curSum int) [][]int {
	curSum += root.Val
	if curSum == sum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	var res [][]int
	if root.Left != nil {
		for _, d := range dfs1(root.Left, sum, curSum) {
			res = append(res, append([]int{root.Val}, d...))
		}
	}

	if root.Right != nil {
		for _, d := range dfs1(root.Right, sum, curSum) {
			res = append(res, append([]int{root.Val}, d...))
		}
	}

	return res
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	return dfs1(root, sum, 0)
}
