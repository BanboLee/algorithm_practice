package sliding_window

/*
leetcode 1074. 元素和为目标值的子矩阵数量
链接：https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target

给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。
子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。
如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。

示例 1：
输入：matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
输出：4
解释：四个只含 0 的 1x1 子矩阵。

示例 2：
输入：matrix = [[1,-1],[-1,1]], target = 0
输出：5
解释：两个 1x2 子矩阵，加上两个 2x1 子矩阵，再加上一个 2x2 子矩阵。

提示：
1 <= matrix.length <= 300
1 <= matrix[0].length <= 300
-1000 <= matrix[i] <= 1000
-10^8 <= target <= 10^8

*/

/*
   在获得每一行每个位置的前缀和之后，对列进行枚举，获得两个列，
   此时滑动行，相当于遍历两列之间所有的矩阵
*/
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var (
		rawLen = len(matrix)
		colLen = len(matrix[0])
		rawSum = make([][]int, rawLen)
		res    int
	)
	// 获得每个位置的行前缀和
	for i := range matrix {
		rawSum[i] = make([]int, colLen)
		rawSum[i][0] = matrix[i][0]
		for j := 1; j < colLen; j++ {
			rawSum[i][j] = rawSum[i][j-1] + matrix[i][j]
		}
	}

	// 固定第一个列
	for i := 0; i < colLen; i++ {
		// 固定第二个列
		for j := i; j < colLen; j++ {
			// 此时将宽度固定，看成一个竖起来的一维数组处理
			var mp = make(map[int]int)
			mp[0] = 1
			curSum := 0
			// 对行进行滑动窗口操作，寻找值为target的矩阵
			for k := 0; k < rawLen; k++ {
				// 对宽度的行进行入窗, 要减去左边窗外的前缀和
				if i > 0 {
					curSum += rawSum[k][j] - rawSum[k][i-1]
				} else {
					curSum += rawSum[k][j]
				}

				// 符合条件
				res += mp[curSum-target]
				mp[curSum]++
			}
		}
	}

	return res
}
