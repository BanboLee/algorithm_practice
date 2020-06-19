package tree

/*
leetcode 617. 合并二叉树
链接：https://leetcode-cn.com/problems/merge-two-binary-trees

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，
那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

示例 1:
输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
注意: 合并必须从两个树的根节点开始。
*/

// 递归
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// 非递归
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	stack := [][]*TreeNode{{t1, t2}}

	for len(stack) != 0 {
		cur := stack[0]
		stack = stack[1:]

		if cur[0] == nil || cur[1] == nil {
			continue
		}

		cur[0].Val += cur[1].Val
		if cur[0].Left == nil {
			cur[0].Left = cur[1].Left
		} else {
			stack = append(stack, []*TreeNode{cur[0].Left, cur[1].Left})
		}

		if cur[0].Right == nil {
			cur[0].Right = cur[1].Right
		} else {
			stack = append(stack, []*TreeNode{cur[0].Right, cur[1].Right})
		}

	}
	return t1
}
