package graph

/*
leetcode 200. 岛屿数量
链接：https://leetcode-cn.com/problems/number-of-islands

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
示例 1:
输入:
11110
11010
11000
00000
输出: 1
示例 2:
输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
*/

// 方法一： 深度优先搜索
func dfs(r, c int, grid [][]byte) {
	rawNum := len(grid)
	colNum := len(grid[0])

	// 检查过的置为0
	grid[r][c] = '0'

	if r-1 >= 0 && grid[r-1][c] == '1' {
		dfs(r-1, c, grid)
	}

	if r+1 < rawNum && grid[r+1][c] == '1' {
		dfs(r+1, c, grid)
	}

	if c-1 >= 0 && grid[r][c-1] == '1' {
		dfs(r, c-1, grid)
	}

	if c+1 < colNum && grid[r][c+1] == '1' {
		dfs(r, c+1, grid)
	}

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				res++
				dfs(r, c, grid)
			}
		}
	}

	return res
}

type unionFindSet struct {
	Count  int
	parent []int
	rank   []int
}

func (ufs *unionFindSet) init(grid [][]byte) {
	n := len(grid[0])
	ufs.parent = make([]int, n*len(grid))
	ufs.rank = make([]int, n*len(grid))

	for i, raw := range grid {
		for j, cell := range raw {
			if cell == '1' {
				ufs.parent[n*i+j] = n*i + j
				ufs.Count++
			} else {
				ufs.parent[n*i+j] = -1
			}
		}
	}
}

func (ufs *unionFindSet) find(x int) int {
	if x != ufs.parent[x] {
		// 查找root并路径压缩
		ufs.parent[x] = ufs.find(ufs.parent[x])
	}
	return ufs.parent[x]
}

func (ufs *unionFindSet) union(x, y int) {
	x, y = ufs.find(x), ufs.find(y)
	// root不同才要合并
	if x != y {
		// 有一个陆地被合并
		ufs.Count--
		if ufs.rank[x] < ufs.rank[y] {
			x, y = y, x
		}
		ufs.parent[y] = x
		if ufs.rank[x] == ufs.rank[y] {
			ufs.rank[x]++
		}
	}
}

func numIslands1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ufs := &unionFindSet{}
	ufs.init(grid)

	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					ufs.union(n*i+j, n*(i-1)+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					ufs.union(n*i+j, n*(i+1)+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					ufs.union(n*i+j, n*i+j-1)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					ufs.union(n*i+j, n*i+j+1)
				}
			}
		}
	}

	return ufs.Count
}
