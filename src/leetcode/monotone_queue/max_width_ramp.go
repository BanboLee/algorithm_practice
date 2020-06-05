package monotone_queue

import "fmt"

/*
leetcode 962. 最大宽度坡
链接：https://leetcode-cn.com/problems/maximum-width-ramp

给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
找出 A 中的坡的最大宽度，如果不存在，返回 0 。

示例 1：
输入：[6,0,8,2,1,5]
输出：4
解释：
最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.

示例 2：
输入：[9,8,1,0,1,9,4,0,4,1]
输出：7
解释：
最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.

提示：
2 <= A.length <= 50000
0 <= A[i] <= 50000
*/

/*
   为何使用单调递减栈: 由题目可知，结果必然是以这个栈中的某个i为坡底的
   反证：如果某个k不在该递减栈中，并且有j-k为最大坡，那么必然有k小于其之前所有的元素，那么必然会被加入该栈
   求结果: 既然得到解的序列，反向遍历数组，依次和栈顶元素比较更新结果
*/
func maxWidthRamp(A []int) int {
	var (
		res   int
		stack []int
	)

	// step1. 构建单调递减栈
	for i, v := range A {
		if stack == nil || A[stack[len(stack)-1]] > v {
			stack = append(stack, i)
		}
	}

	fmt.Printf("stack=%v\n", stack)
	// step2. 反向遍历求结果
	for i := len(A) - 1; i >= 0 && len(stack) != 0; i-- {
		for len(stack) != 0 && A[stack[len(stack)-1]] <= A[i] {
			res = Max(res, i-stack[len(stack)-1])
			// 如果当前栈顶已经符合要求，那么立即出栈，因为这已经是它目前可以形成的最大坡
			stack = stack[:len(stack)-1]
		}
	}

	return res
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
