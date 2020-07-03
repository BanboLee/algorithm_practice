package dict_tree

/*
leetcode 677. 键值映射
链接：https://leetcode-cn.com/problems/map-sum-pairs

实现一个 MapSum 类里的两个方法，insert 和 sum。
对于方法 insert，你将得到一对（字符串，整数）的键值对。
字符串表示键，整数表示值。如果键已经存在，那么原来的键值对将被替代成新的键值对。
对于方法 sum，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。

示例 1:
输入: insert("apple", 3), 输出: Null
输入: sum("ap"), 输出: 3
输入: insert("app", 2), 输出: Null
输入: sum("ap"), 输出: 5
*/

type MapSum struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor1() MapSum {
	return MapSum{&TrieNode{}}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	for i := 0; i < len(key); i++ {
		if cur.children[key[i]-'a'] == nil {
			cur.children[key[i]-'a'] = &TrieNode{}
		}
		cur = cur.children[key[i]-'a']
	}
	cur.isEnd = true
	cur.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		if cur.children[prefix[i]-'a'] == nil {
			return 0
		}
		cur = cur.children[prefix[i]-'a']
	}

	return cur.getSum()
}
