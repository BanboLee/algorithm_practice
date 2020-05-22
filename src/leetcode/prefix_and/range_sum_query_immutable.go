package prefix_and

/*
leetcode 303. 区域和检索 - 数组不可变
链接：https://leetcode-cn.com/problems/range-sum-query-immutable

给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：
给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

说明:
你可以假设数组不可变。
会多次调用 sumRange 方法。
*/

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{[]int{0}}

	for i, v := range nums {
		na.prefixSum = append(na.prefixSum, na.prefixSum[i]+v)
	}
	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.prefixSum[j+1] - this.prefixSum[i]
}

/*
leetcode 304. 二维区域和检索 - 矩阵不可变
链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable

给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。
上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。

示例:
给定 matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12

说明:
你可以假设矩阵不可变。
会多次调用 sumRegion 方法。
你可以假设 row1 ≤ row2 且 col1 ≤ col2。

分析：
    使用前缀和方法：假设所求矩阵四个角为a,b,c,d, 原点为o, 使用中间矩阵存储其左上角的和，则
                  sum(a, d) = sum(o, d) - sum(o, b) - sum(o, c) + sum(o, a)
*/

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	nm := NumMatrix{make([][]int, len(matrix))}
	// 计算前缀和
	for i := range nm.sumMatrix {
		for j, v := range matrix[i] {
			// left, up, left-up
			left, up, lu := 0, 0, 0
			if i != 0 {
				up = nm.sumMatrix[i-1][j]
			}
			if j != 0 {
				left = nm.sumMatrix[i][j-1]
			}
			if i != 0 && j != 0 {
				lu = nm.sumMatrix[i-1][j-1]
			}
			nm.sumMatrix[i] = append(nm.sumMatrix[i], up+left-lu+v)
		}
	}
	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 通过前缀和计算结果
	left, up, lu := 0, 0, 0
	if row1 != 0 {
		up = this.sumMatrix[row1-1][col2]
	}
	if col1 != 0 {
		left = this.sumMatrix[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		lu = this.sumMatrix[row1-1][col1-1]
	}
	return this.sumMatrix[row2][col2] - left - up + lu
}
