package tree

/*
leetcode 199. 二叉树的右视图
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view

给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res []int

	for len(queue) != 0 {
		l := len(queue)
		res = append(res, queue[l-1].Val)
		for i := 0; i < l; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[l:]
	}
	return res
}
