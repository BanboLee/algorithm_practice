package prefix_and

/*
leetcode 1124. 表现良好的最长时间段
链接：https://leetcode-cn.com/problems/longest-well-performing-interval

给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
请你返回「表现良好时间段」的最大长度。

示例 1：
输入：hours = [9,9,6,0,6,6,9]
输出：3
解释：最长的表现良好时间段是 [9,9,6]。

提示：
1 <= hours.length <= 10000
0 <= hours[i] <= 16
*/

/*
   差值 + 滑动窗口
*/
func longestWPI(hours []int) int {
	for i := range hours {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = 0
		}
	}

	l, r, res, preSum := 0, 0, 0, 0
	for ; r < len(hours); r++ {
		// 入窗
		preSum += hours[r]

		// 出窗
		for l < r && preSum*2 <= r-l+1 {
			preSum -= hours[l]
			l++
		}

		if l < r && res < r-l+1 {
			res = r - l + 1
		}
	}

	return res
}
