package graph

/*
leetcode 79. 单词搜索
链接：https://leetcode-cn.com/problems/word-search

给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：
board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/

func dfs2(r, c int, word string, board [][]byte, mark [][]bool) bool {
	if word[0] != board[r][c] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	mark[r][c] = true
	flag := false
	if r-1 >= 0 && !mark[r-1][c] {
		flag = flag || dfs2(r-1, c, word[1:], board, mark)
	}

	if r+1 < len(board) && !mark[r+1][c] {
		flag = flag || dfs2(r+1, c, word[1:], board, mark)
	}

	if c+1 < len(board[r]) && !mark[r][c+1] {
		flag = flag || dfs2(r, c+1, word[1:], board, mark)
	}

	if c-1 >= 0 && !mark[r][c-1] {
		flag = flag || dfs2(r, c-1, word[1:], board, mark)
	}

	mark[r][c] = false
	return flag
}

func exist(board [][]byte, word string) bool {
	mark := make([][]bool, len(board))
	for i := range mark {
		mark[i] = make([]bool, len(board[i]))
	}
	for r, raw := range board {
		for c := range raw {
			if dfs2(r, c, word, board, mark) {
				return true
			}
		}
	}

	return false
}
