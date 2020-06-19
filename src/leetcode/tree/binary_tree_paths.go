package tree

import (
	"fmt"
	"strconv"
)

/*
leetcode 257. 二叉树的所有路径
链接：https://leetcode-cn.com/problems/binary-tree-paths

给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
输入:
   1
 /   \
2     3
 \
  5
输出: ["1->2->5", "1->3"]
解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	var res []string
	if root.Left != nil {
		for _, s := range binaryTreePaths(root.Left) {
			res = append(res, fmt.Sprintf("%d->%s", root.Val, s))
		}
	}
	if root.Right != nil {
		for _, s := range binaryTreePaths(root.Right) {
			res = append(res, fmt.Sprintf("%d->%s", root.Val, s))
		}
	}

	return res
}
