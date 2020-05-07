package sliding_window

import "math"

/*
leet-code 209 长度最小的子数组
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum

给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。
如果不存在符合条件的连续子数组，返回 0。

示例:
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。

进阶:
如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

分析：
一、滑动窗口(双指针) O(n)
二、二分查找 O(nlogn)
*/

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLen, l, r, curSum := math.MaxInt32, 0, 0, 0
	for ; r < len(nums); r++ {
		curSum += nums[r]
		for curSum >= s {
			if minLen > r-l+1 {
				minLen = r - l + 1
			}
			curSum -= nums[l]
			l++
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}

// 二分查找法
func minSubArrayLen1(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 1. 定义数组存储之前数字之和
	var sums = make([]int, len(nums)+1)
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}

	minLen := math.MaxInt32
	// 2. 遍历每个数字
	for i, v := range nums {
		// 3. 减去当前值，得到需要除当前值外的最大和
		s1 := s - v
		// 4. 从当前值开始，向右寻找，直到符合条件
		/* FIXME:
				    在二分查找中，这里应该去sums中找到第一个>=s+sums[i - 1]的值bound，
		            然后得到i到bound最小值bound-i+1
		*/
		for j := i; j < len(nums); j++ {
			if sums[j]-sums[i] >= s1 {
				if minLen > j-i+1 {
					minLen = j - i + 1
				}
				// 找到第一个即可
				break
			}
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}
