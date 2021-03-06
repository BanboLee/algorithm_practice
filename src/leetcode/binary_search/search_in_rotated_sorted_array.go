package binary_search

/*
leetcode 33. 搜索旋转排序数组
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array

假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。

示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1

	/*
	   旋转之后都是升序序列, 并且前半段序列比后半段序列起点高
	   如果target<nums[0], 说明结果需要到旋转点的后半段查找
	   如果target>=nums[0], 说明结果需要到旋转点的前半段查找
	*/
	for l <= r {
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] < nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}
