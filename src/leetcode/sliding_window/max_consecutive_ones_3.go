package sliding_window

/*
leetcode 1004. 最大连续1的个数 III
链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii

给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
返回仅包含 1 的最长（连续）子数组的长度。

示例 1：
输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。

示例 2：
输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。

提示：
1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为 0 或 1
*/

/*
   使用滑动窗口扫描，
   1. 如果右边为0，并且还有变为1的名额，名额-1，
   2. 如果右边为0，并且没有名额了，从左边出窗一个(注意名额变化)，后边入窗一个
   3. 如果右边为1， 直接入窗
*/
func longestOnes(A []int, K int) int {
	l, r, flipped, res := 0, 0, 0, 0

	for ; r < len(A); r++ {
		// 处理好出窗
		if A[r] == 0 {
			flipped++
		}
		for flipped > K {
			if A[l] == 0 {
				flipped--
			}
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
