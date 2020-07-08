package binary_search

/*
leetcode 153. 寻找旋转排序数组中的最小值
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array

假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
请找出其中最小的元素。
你可以假设数组中不存在重复元素。
示例 1:
输入: [3,4,5,1,2]
输出: 1
示例 2:
输入: [4,5,6,7,0,1,2]
输出: 0
*/

// 找到变化的那个点，确定最小值
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	// 从小到大排序并被打乱，那么中间必定有个值比第一个值小
	for l <= r {
		mid := l + (r-l)/2
		if mid == 0 {
			// 此时有可能mid=0 [2,1]
			if nums[1] < nums[0] {
				return nums[1]
			}
			return nums[0]
		} else if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			// 结果在左边, 判断这个值是不是开始变化那个点
			if mid == 0 || nums[mid-1] > nums[mid] {
				return nums[mid]
			} else {
				r = mid - 1
			}
		}
	}
	return nums[0]
}
