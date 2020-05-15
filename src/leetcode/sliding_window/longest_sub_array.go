package sliding_window

/*
leetcode 1438. 绝对差不超过限制的最长连续子数组
链接：https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit

给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
如果不存在满足条件的子数组，则返回 0 。

示例 1：
输入：nums = [8,2,4,7], limit = 4
输出：2
解释：所有子数组如下：
[8] 最大绝对差 |8-8| = 0 <= 4.
[8,2] 最大绝对差 |8-2| = 6 > 4.
[8,2,4] 最大绝对差 |8-2| = 6 > 4.
[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
[2] 最大绝对差 |2-2| = 0 <= 4.
[2,4] 最大绝对差 |2-4| = 2 <= 4.
[2,4,7] 最大绝对差 |2-7| = 5 > 4.
[4] 最大绝对差 |4-4| = 0 <= 4.
[4,7] 最大绝对差 |4-7| = 3 <= 4.
[7] 最大绝对差 |7-7| = 0 <= 4.
因此，满足题意的最长子数组的长度为 2 。

示例 2：
输入：nums = [10,1,2,4,7,2], limit = 5
输出：4
解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。

示例 3：
输入：nums = [4,2,2,2,4,4,2,2], limit = 0
输出：3

提示：
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
0 <= limit <= 10^9
*/

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

/*
   单调队列+滑动窗口
   使用两个单调队列保存当前窗口的最大绝对值和最小绝对值的索引，
   每次滑动去更新这个窗口。
*/
func longestSubarray(nums []int, limit int) int {
	if len(nums) == 1 {
		return 1
	}

	l, r, res := 0, 0, 0
	dqMax, dqMin := make([]int, 0), make([]int, 0)
	for ; r < len(nums); r++ {
		// 入窗
		// 新来的大于之前队列尾，就让队尾出队，让其进入
		for len(dqMax) > 0 && abs(nums[r]) > abs(nums[dqMax[len(dqMax)-1]]) {
			dqMax = dqMax[:len(dqMax)-1]
		}
		if len(dqMax) == 0 || abs(nums[r]) <= abs(nums[dqMax[len(dqMax)-1]]) {
			dqMax = append(dqMax, r)
		}
		for len(dqMin) > 0 && abs(nums[r]) < abs(nums[dqMin[len(dqMin)-1]]) {
			dqMin = dqMin[:len(dqMin)-1]
		}
		if len(dqMin) == 0 || abs(nums[r]) >= abs(nums[dqMin[len(dqMin)-1]]) {
			dqMin = append(dqMin, r)
		}

		// 出窗, 不满足条件就从左边出窗
		for abs(nums[dqMax[0]])-abs(nums[dqMin[0]]) > limit {
			if l == dqMax[0] {
				dqMax = dqMax[1:]
			} else if l == dqMin[0] {
				dqMin = dqMin[1:]
			}
			l++
		}

		// 获取结果
		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}
