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
	queue []*TreeNode
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[null]"
	}
	this.queue = []*TreeNode{root}
	s := "["

	for len(this.queue) != 0 {
		cur := this.queue[0]
		this.queue = this.queue[1:]
		if cur == nil {
			s += "null,"
			continue
		} else {
			s += strconv.Itoa(cur.Val)
		}
		s += ","

		this.queue = append(this.queue, cur.Left)
		this.queue = append(this.queue, cur.Right)
	}

	s = s[:len(s)-1]
	s += "]"
	fmt.Println(s)
	return s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1]
	if len(data) == 0 {
		return nil
	}

	nodes := strings.Split(data, ",")
	this.queue = make([]*TreeNode, 0)
	for _, node := range nodes {
		if node == "null" {
			this.queue = append(this.queue, nil)
		} else {
			val, _ := strconv.Atoi(node)
			this.queue = append(this.queue, &TreeNode{val, nil, nil})
		}
	}

	for i, node := range this.queue {
		if i*2+1 < len(this.queue) {
			node.Left = this.queue[i*2+1]
		}
		if i*2+2 < len(this.queue) {
			node.Right = this.queue[i*2+2]
		}
	}

	return this.queue[0]
}
