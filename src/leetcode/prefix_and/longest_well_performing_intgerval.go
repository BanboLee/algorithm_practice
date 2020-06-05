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
	1. 状态压缩，转换为1和-1表示
	2. 前缀和, 从开始到当前位置是否是良好
	3. 转化为坡度问题，实际是求sum(i, j) > 0 时的最长宽度坡(i, j)
	个人认为难点就在怎么想到从前缀和转化为这个形式的最长宽度坡问题
*/
func longestWPI(hours []int) int {
	if len(hours) == 0 {
		return 0
	}
	// step1. 表示数据状态
	for i := range hours {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	// step2. 计算前缀和
	var preSum = make([]int, len(hours)+1)
	preSum[0] = 0
	for i := 1; i < len(hours)+1; i++ {
		preSum[i] = preSum[i-1] + hours[i-1]
	}

	// step3. 坡度问题
	// 3.1 构造单调栈
	var stack []int
	for i := 0; i < len(preSum); i++ {
		if len(stack) == 0 || preSum[stack[len(stack)-1]] > preSum[i] {
			stack = append(stack, i)
		}
	}

	// 3.2 求解
	var res = 0
	// 如果 j <= res了，结果不会再更新为更大了
	for j := len(preSum) - 1; j > res && len(stack) != 0; j-- {
		for len(stack) != 0 && preSum[j] > preSum[stack[len(stack)-1]] {
			// 保证sum(i, j) > 0, 其实是(i, j]的距离不然就是j-i+1
			res = max(res, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return res
}

/*
	来自大佬的题解：https://leetcode-cn.com/problems/longest-well-performing-interval/solution/bie-gen-lao-fu-ti-shi-yao-dan-diao-zhan-by-li-zi-h/
*/
func longestWPI1(hours []int) int {
	var (
		res int
		cur int
		mp  = make(map[int]int)
	)

	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			cur++
		} else {
			cur--
		}

		// 当cur大于0的时候，到目前为止都是良好的
		if cur > 0 {
			res = i + 1
		} else {
			// 这里主要是处理好怎样使用hash结合前缀和
			/*
				用 cur - 1 去字典里找，如果找到了下标j，那么就说明从0到 j 的前缀和是 cur-1，
				而从0到 i 的前缀和是 cur，那么显然从 j 到 i的和是（cur - (cur - 1)） = 1 > 0，
				也就是说从 j+1到 i 的表现肯定是满足的，并且由于 j 是 cur-1中最小的，所以 i-j 是最大的。
			*/
			if _, ok := mp[cur-1]; ok {
				// 如果之前出现过这个值，那么第一次这个值出现到现在，一定是良好的
				res = max(res, i-mp[cur-1])
			} else {
				mp[cur] = i
			}
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
