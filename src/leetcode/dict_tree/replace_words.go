package dict_tree

import "strings"

/*
leetcode 648. 单词替换
链接：https://leetcode-cn.com/problems/replace-words

在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。
例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
现在，给定一个由许多词根组成的词典和一个句子。
你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
你需要输出替换之后的句子。

示例：
输入：dict(词典) = ["cat", "bat", "rat"] sentence(句子) = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

提示：
输入只包含小写字母。
1 <= dict.length <= 1000
1 <= dict[i].length <= 100
1 <= 句中词语数 <= 1000
1 <= 句中词语长度 <= 1000
*/

type TrieNode struct {
	isEnd    bool
	Val      int
	children [26]*TrieNode
}

func (tn *TrieNode) insert(word string) {
	cur := tn
	for i := 0; i < len(word); i++ {
		if cur.children[word[i]-'a'] == nil {
			cur.children[word[i]-'a'] = &TrieNode{}
		}
		cur = cur.children[word[i]-'a']
	}
	cur.isEnd = true
}

func (tn *TrieNode) hasPrefix(word string) string {
	cur := tn
	i := 0
	for i < len(word) {
		cur = cur.children[word[i]-'a']

		if cur == nil {
			return ""
		}

		if cur.isEnd {
			return word[0 : i+1]
		}
		i++
	}

	if cur.isEnd {
		return word[0:i]
	}
	return ""
}

func replaceWords(dict []string, sentence string) string {
	tn := &TrieNode{}

	for _, word := range dict {
		tn.insert(word)
	}

	words := strings.Split(sentence, " ")
	var newWords []string
	for _, word := range words {
		w := tn.hasPrefix(word)
		if w == "" {
			newWords = append(newWords, word)
		} else {
			newWords = append(newWords, w)
		}
	}

	return strings.Join(newWords, " ")
}
