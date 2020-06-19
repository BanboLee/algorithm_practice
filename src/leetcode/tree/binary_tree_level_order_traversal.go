package tree

/*
leetcode 102. 二叉树的层序遍历
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res [][]int

	for len(queue) != 0 {
		l := len(queue)
		var tmp []int
		for i := 0; i < l; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			tmp = append(tmp, cur.Val)
		}
		res = append(res, tmp)
		queue = queue[l:]
	}
	return res
}
