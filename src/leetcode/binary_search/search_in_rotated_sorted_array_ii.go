package binary_search

/*
leetcode 81. 搜索旋转排序数组 II

假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:
输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true

示例 2:
输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false

进阶:
这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
*/

// 跳过无法判定是否有序的位置
func search1(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1

	/*
	   旋转之后都是升序序列, 并且前半段序列比后半段序列起点高
	   如果target<nums[0], 说明结果需要到旋转点的后半段查找
	   如果target>=nums[0], 说明结果需要到旋转点的前半段查找
	*/
	for l <= r {
		mid := (l + r) / 2
		if nums[l] == target || nums[r] == target || nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] {
			l++
			continue
		}
		if nums[l] < nums[mid] {
			if target >= nums[l] && target < nums[mid] {
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
	return false
}
