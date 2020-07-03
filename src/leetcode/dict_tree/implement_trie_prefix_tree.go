package dict_tree

/*
leetcode 208. 实现 Trie (前缀树)
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree

实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
示例:
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:
你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。
*/
func (tn *TrieNode) getSum() int {
	if tn == nil {
		return 0
	}

	sum := 0
	if tn.isEnd {
		sum += tn.Val
	}
	for _, node := range tn.children {
		if node != nil {
			sum += node.getSum()
		}
	}
	return sum
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{&TrieNode{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		if cur.children[word[i]-'a'] == nil {
			cur.children[word[i]-'a'] = &TrieNode{}
		}
		cur = cur.children[word[i]-'a']
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		if cur.children[word[i]-'a'] == nil {
			return false
		}
		cur = cur.children[word[i]-'a']
	}
	if cur.isEnd {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		if cur.children[prefix[i]-'a'] == nil {
			return false
		}
		cur = cur.children[prefix[i]-'a']
	}
	return true
}
