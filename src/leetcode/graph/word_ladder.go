package graph

/*
leetcode 127. 单词接龙
链接：https://leetcode-cn.com/problems/word-ladder

给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:
如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]
输出: 5
解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。

示例 2:
输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]
输出: 0
解释: endWord "cog" 不在字典中，所以无法进行转换。
*/

/*
官方解析1:
想法
利用广度优先搜索搜索从 beginWord 到 endWord 的路径。
算法
1. 对给定的 wordList 做预处理，找出所有的通用状态。将通用状态记录在字典中，键是通用状态，值是所有具有通用状态的单词。
2. 将包含 beginWord 和 1 的元组放入队列中，1 代表节点的层次。我们需要返回 endWord 的层次也就是从 beginWord 出发的最短距离。
3. 为了防止出现环，使用访问数组记录。
4. 当队列中有元素的时候，取出第一个元素，记为 current_word。
5. 找到 current_word 的所有通用状态，并检查这些通用状态是否存在其它单词的映射，这一步通过检查 all_combo_dict 来实现。
6. 从 all_combo_dict 获得的所有单词，都和 current_word 共有一个通用状态，所以都和 current_word 相连，因此将他们加入到队列中。
7. 对于新获得的所有单词，向队列中加入元素 (word, level + 1) 其中 level 是 current_word 的层次。
8. 最终当你到达期望的单词，对应的层次就是最短变换序列的长度。
标准广度优先搜索的终止条件就是找到结束单词。

链接：https://leetcode-cn.com/problems/word-ladder/solution/dan-ci-jie-long-by-leetcode/
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	allComboDict := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			gs := word[:i] + "_" + word[i+1:]
			allComboDict[gs] = append(allComboDict[gs], word)
		}
	}

	type pair struct {
		word  string
		level int
	}

	queue := make([]pair, 0)
	visited := make(map[string]bool)

	queue = append(queue, pair{beginWord, 1})

	for len(queue) != 0 {
		curWord := queue[0].word
		curLevel := queue[0].level
		if curWord == endWord {
			return curLevel
		}
		queue = queue[1:]

		if visited[curWord] {
			continue
		}
		visited[curWord] = true

		for i := 0; i < len(curWord); i++ {
			gs := curWord[:i] + "_" + curWord[i+1:]
			for _, word := range allComboDict[gs] {
				queue = append(queue, pair{word, curLevel + 1})
			}
		}
	}
	return 0
}

// 官方题解2：对向广度优先搜索
type pair struct {
	word  string
	level int
}

func visitNode(queue *[]pair, visited, otherVisited map[string]int, allCombWord map[string][]string) int {
	if len(*queue) == 0 {
		return -1
	}

	node := (*queue)[0]
	*queue = (*queue)[1:]

	curWord := node.word
	curLevel := node.level

	for i := 0; i < len(curWord); i++ {
		gs := curWord[:i] + "_" + curWord[i+1:]

		for _, word := range allCombWord[gs] {
			if otherVisited[word] > 0 {
				return otherVisited[word] + curLevel
			}

			if visited[word] == 0 {
				visited[word] = curLevel + 1
				*queue = append(*queue, pair{word, curLevel + 1})
			}
		}
	}

	return -1

}

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	flag := false
	allCombWord := make(map[string][]string)

	for _, word := range wordList {
		if word == endWord {
			flag = true
		}
		for i := 0; i < len(word); i++ {
			gs := word[:i] + "_" + word[i+1:]
			allCombWord[gs] = append(allCombWord[gs], word)
		}
	}

	if !flag {
		return 0
	}

	queueBegin := []pair{{beginWord, 1}}
	queueEnd := []pair{{endWord, 1}}

	visitedBegin := make(map[string]int)
	visitedEnd := make(map[string]int)
	visitedBegin[beginWord] = 1
	visitedEnd[endWord] = 1

	for len(queueBegin) != 0 && len(queueEnd) != 0 {
		ans := visitNode(&queueBegin, visitedBegin, visitedEnd, allCombWord)
		if ans > -1 {
			return ans
		}

		ans = visitNode(&queueEnd, visitedEnd, visitedBegin, allCombWord)
		if ans > -1 {
			return ans
		}
	}

	return 0
}
