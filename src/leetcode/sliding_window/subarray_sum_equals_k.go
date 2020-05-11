package sliding_window

/*
leetcode 560. 和为K的子数组
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

说明 :
数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

分析：
一. 滑动窗口 O(n²)
二. 前缀和+hash O(n)
	1. 使用hash表示前缀和出现次数
	2. 初始化时，要确认0出现了一次
	3. 前缀和的意义在于，当你发现此时sum-k已经出现过时，
       说明对应的前缀和到当前值的和为一个k，
       假如之前出现过n个sum-k，那么就可以和当前值组成n个和为k的数组。
*/

func subarraySum(nums []int, k int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		l := i
		sum := 0
		for r := l; r < len(nums); r++ {
			sum += nums[r]
			if sum == k {
				res++
			}
		}
	}

	return res
}

func subarraySumWithPreSum(nums []int, k int) int {
	var (
		res    int
		curSum int
		mp     = make(map[int]int, 0)
	)

	// 初始时候，前缀和为0，所以字典中0出现了一次
	mp[0] = 1
	for _, v := range nums {
		curSum += v
		//  若当有前缀和+k能得到当前值, 则该前缀和位置到当前位置为一个有效数组
		if mp[curSum-k] > 0 {
			res += mp[curSum-k]
		}
		mp[curSum]++
	}

	return res
}
