package tree

/*
leetcode 105. 从前序与中序遍历序列构造二叉树
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// get mid
	t := &TreeNode{preorder[0], nil, nil}
	for i, v := range inorder {
		if v == t.Val {
			t.Left = buildTree(preorder[1:1+len(inorder[:i])], inorder[:i])
			t.Right = buildTree(preorder[1+len(inorder[:i]):], inorder[i+1:])
			break
		}
	}

	return t
}
