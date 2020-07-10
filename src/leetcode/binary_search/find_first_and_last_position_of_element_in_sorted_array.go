package binary_search

/*
leetcode 34. 在排序数组中查找元素的第一个和最后一个位置
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

// 分两部分进行，找到左边边界，再找到右边界
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := searchMargin(nums, target, true)
	// 判断左边界是否正确
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	return []int{left, searchMargin(nums, target, false) - 1}
}

func searchMargin(nums []int, target int, left bool) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > target || (left && target == nums[mid]) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
