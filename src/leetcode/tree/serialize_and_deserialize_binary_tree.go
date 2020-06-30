package tree

import (
	"fmt"
	"strconv"
	"strings"
)

/*
leetcode 297. 二叉树的序列化与反序列化
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree

序列化是将一个数据结构或者对象转换为连续的比特位的操作，
进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:
你可以将以下二叉树：
    1
   / \
  2   3
     / \
    4   5
序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
*/

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

type Codec struct {
	queue   []*TreeNode
	nodeVal []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// 选择一种遍历方式即可，不要用广度优先遍历，反序列化的时候会坑，最好是先序遍历
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := "["
	this.queue = []*TreeNode{root}
	for len(this.queue) != 0 {
		cur := this.queue[len(this.queue)-1]
		this.queue = this.queue[:len(this.queue)-1]
		if cur == nil {
			res += "null,"
			continue
		} else {
			res += strconv.Itoa(cur.Val) + ","
		}
		// 使用栈， 按先序，先压入右子树，再压如左子树
		this.queue = append(this.queue, cur.Right)
		this.queue = append(this.queue, cur.Left)
	}

	return res[:len(res)-1] + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	data = data[1 : len(data)-1]
	if data == "" {
		return nil
	}
	this.nodeVal = strings.Split(data, ",")
	fmt.Printf("data=%v\n", this.nodeVal)
	return this.redeserialize()

}

// 递归
func (this *Codec) redeserialize() *TreeNode {
	if this.nodeVal[0] == "null" {
		this.nodeVal = this.nodeVal[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.nodeVal[0])
	root := &TreeNode{val, nil, nil}
	this.nodeVal = this.nodeVal[1:]
	root.Left = this.redeserialize()
	root.Right = this.redeserialize()
	return root
}
