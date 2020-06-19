package tree

/*
leetcode 103. 二叉树的锯齿形层次遍历
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal

给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：
[
  [3],
  [20,9],
  [15,7]
]
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res [][]int
	flag := false

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

			if flag {
				tmp = append([]int{cur.Val}, tmp...)
			} else {
				tmp = append(tmp, cur.Val)
			}
		}

		res = append(res, tmp)
		queue = queue[l:]
		flag = !flag
	}
	return res
}
