package array

import (
	"fmt"
)

// 螺旋矩阵1  按顺时针打印矩阵
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	var result = make([]int, len(matrix)*len(matrix[0]))
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	const (
		UP    = 0
		DOWN  = 1
		LEFT  = 2
		RIGHT = 3
	)
	// 0-up 1-down 2-left 3-right
	flag := RIGHT
	step := 0

	for i := range result {
		switch flag {
		case UP:
			result[i] = matrix[step][left]
			if step == up {
				step = left + 1
				flag = RIGHT
				left++
			} else {
				step--
			}
		case DOWN:
			result[i] = matrix[step][right]
			if step == down {
				step = right - 1
				flag = LEFT
				right--
			} else {
				step++
			}
		case RIGHT:
			result[i] = matrix[up][step]
			if step == right {
				step = up + 1
				flag = DOWN
				up++
			} else {
				step++
			}
		case LEFT:
			result[i] = matrix[down][step]
			if step == left {
				step = down - 1
				flag = UP
				down--
			} else {
				step--
			}
		}
	}
	return result
}

// 螺旋矩阵2 按顺时针填充生成矩阵
func GenerateMatrix(n int) [][]int {
	var total = n * n
	var result = make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	up, down, left, right := 0, n-1, 0, n-1
	const (
		UP    = 0
		DOWN  = 1
		LEFT  = 2
		RIGHT = 3
	)
	// 0-up 1-down 2-left 3-right
	flag := RIGHT
	step := 0

	for i := 0; i < total; i++ {
		switch flag {
		case UP:
			result[step][left] = i + 1
			if step == up {
				flag = RIGHT
				left++
				step = left
			} else {
				step--
			}
		case DOWN:
			result[step][right] = i + 1
			if step == down {
				flag = LEFT
				right--
				step = right
			} else {
				step++
			}
		case LEFT:
			result[down][step] = i + 1
			if step == left {
				flag = UP
				down--
				step = down
			} else {
				step--
			}
		case RIGHT:
			result[up][step] = i + 1
			if step == right {
				flag = DOWN
				up++
				step = up
			} else {
				step++
			}
		}
	}
	return result
}

/*
原地矩阵旋转90度
给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

分析：
1. 可以转置然后翻转每一行
2. 一层一层转进去
3. 按矩形转
*/
func Rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	// 一圈一圈转
	// bound
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	for up < down {
		for i := 0; i < right-left; i++ {
			x, y, lx, ly := up, left+i, right, down
			tmp := matrix[x][y]
			x1, y1 := lx-y+up, x
			matrix[x][y] = matrix[x1][y1]
			x2, y2 := ly-y1+up, x1
			matrix[x1][y1] = matrix[x2][y2]
			x3, y3 := lx-y2+up, x2
			matrix[x2][y2] = matrix[x3][y3]
			matrix[x3][y3] = tmp
		}
		up++
		down--
		left++
		right--
	}
}

/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

进阶:
一个直接的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？

分析:
1. 使用一个新的矩阵，直接生成
2. 使用两个数组分别表示行和列是否要置0
3. 标记每行的第一和每列的第一个元素
*/
func SetZeroes(matrix [][]int) {
	raw, col := make([]int, len(matrix)), make([]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				raw[i] = -1
				col[j] = -1
			}
		}
	}
	fmt.Printf("raw=%v\ncol=%v\n", raw, col)
	for i := range matrix {
		for j := range matrix[i] {
			if raw[i] == -1 || col[j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}

/*
除自身以外数组的乘积
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

输入: [1,2,3,4]
输出: [24,12,8,6]

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

分析：
output[i] = nums[0] * nums[1] * ... * nums[i - 1] * nums[i + 1] * ... * nums[n - 1] = prev[i] * back[i]
prev[i] = nums[0] * nums[1] * ... * nums[i - 1] = prev[i - 1] * nums[i - 1]
back[i] = nums[i + 1] * ... * nums[n - 1] = back[i + 1] * nums[i + 1]
时间复杂度 O(n)
空间复杂度 O(n)

分析：
1. 先从左边乘过去，将第一遍结果放到输出数组中，再从右边乘回来
2. 两边同时累乘

*/
func ProductExceptSelf(nums []int) []int {
	prev := make([]int, len(nums))
	prev[0] = 1
	for i := 1; i < len(nums); i++ {
		prev[i] = prev[i-1] * nums[i-1]
	}

	back := 1
	for j := len(nums) - 1; j >= 0; j-- {
		prev[j] = prev[j] * back
		back = back * nums[j+1]
	}
	return prev
}
func ProductExceptSelf2(nums []int) []int {
	output := make([]int, len(nums))
	for i := range output {
		output[i] = 1
	}
	prev, next := 1, 1
	for i := range nums {
		output[i] *= prev
		prev *= nums[i]
		output[len(nums)-1-i] *= next
		next *= nums[len(nums)-1-i]
	}
	return output
}

/*
搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。

分析：
时间复杂度: O(m + n)
选左上角，往右走和往下走都增大，不能选
选右下角，往上走和往左走都减小，不能选
选左下角，往右走增大，往上走减小，可选
选右上角，往下走增大，往左走减小，可选
*/

func SearchMatrix(matrix [][]int, target int) bool {
	raw, col := 0, len(matrix[0])-1
	for col >= 0 && raw < len(matrix) {
		if matrix[raw][col] == target {
			return true
		} else if matrix[raw][col] > target {
			col--
		} else {
			raw++
		}
	}
	return false
}
