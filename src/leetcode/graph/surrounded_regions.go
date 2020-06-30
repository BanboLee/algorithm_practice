package graph

/*
leetcode 130. 被围绕的区域
链接：https://leetcode-cn.com/problems/surrounded-regions

给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:
X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：
X X X X
X X X X
X X X X
X O X X
解释:
被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/

func dfs1(r, c int, board [][]byte) {
	board[r][c] = '_'

	if r-1 >= 0 && board[r-1][c] == 'O' {
		dfs1(r-1, c, board)
	}
	if r+1 < len(board) && board[r+1][c] == 'O' {
		dfs1(r+1, c, board)
	}
	if c-1 >= 0 && board[r][c-1] == 'O' {
		dfs1(r, c-1, board)
	}
	if c+1 < len(board[r]) && board[r][c+1] == 'O' {
		dfs1(r, c+1, board)
	}
}

// 先按岛屿数量那种方式，将连接边界的全都改为其他字符
func solve(board [][]byte) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if (r == 0 || c == 0 || r == len(board)-1 || c == len(board[r])-1) && board[r][c] == 'O' {
				dfs1(r, c, board)
			}
		}
	}

	for r, raw := range board {
		for c, col := range raw {
			if col == 'O' {
				board[r][c] = 'X'
			}
			if col == '_' {
				board[r][c] = 'O'
			}
		}
	}
}
