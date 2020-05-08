package sliding_window

/*
leetcode 992. K 个不同整数的子数组
链接：https://leetcode-cn.com/problems/subarrays-with-k-different-integers

给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。
（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
返回 A 中好子数组的数目。

示例 1：
输入：A = [1,2,1,2,3], K = 2
输出：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
示例 2：
输入：A = [1,2,1,3,4], K = 3
输出：3
解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].

提示：
1 <= A.length <= 20000
1 <= A[i] <= A.length
1 <= K <= A.length

*/

func subarraysWithKDistinct(A []int, K int) int {
	var res, l int
	var table = make(map[int]int)
	for i := 0; i < len(A); i++ {
		// 入窗
		table[A[i]]++

		// 出窗
		for len(table) > K {
			table[A[l]]--
			if table[A[l]] == 0 {
				delete(table, A[l])
			}
			l++
		}

		if len(table) == K {
			// 使用双指针+滑动窗口确定所有可能
			var h = make([]int, len(A))
			l1 := l
			for l1 <= i {
				if h[A[l1]-1] == 0 {
					h[A[l1]-1] = table[A[l1]]
				}
				res++
				h[A[l1]-1]--
				if h[A[l1]-1] == 0 {
					break
				}
				l1++
			}
		}
	}
	return res
}
