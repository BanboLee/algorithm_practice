package array

import "sort"

/*
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

说明:
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

分析：
一个指针指向最近一个0，一个指向最近一个非0
*/

func MoveZeroes(nums []int) {
	// i for closest 0, j for closest non 0.
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == 0 {
			if j < i {
				j = i + 1
			}
			for j < len(nums) {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j++
			}
		} else {
			i++
		}
	}
}

/*
判断子序列
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

示例 1:
s = "abc", t = "ahbgdc"
返回 true.
示例 2:
s = "axc", t = "ahbgdc"
返回 false.

后续挑战 :
如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

分析:
两个指针，一个指向子序列，一个指向元序列
*/
func IsSubsequence(s string, t string) bool {
	// i for s, j for t
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if t[j] == s[i] {
			i++
			j++
		} else {
			j++
		}
	}

	if i == len(s) {
		return true
	}
	return false
}

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，
垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
输入：[1,8,6,2,5,4,8,3,7]
输出：49

分析：
	双指针从两头到中间
*/
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	max := min(height[l], height[r]) * (r - l)
	for l <= r {
		if max < min(height[l], height[r])*(r-l) {
			max = min(height[l], height[r]) * (r - l)
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

分析：
1. 固定一个数n，然后双指针找和为-n的
2. 如果和大于-n就右边左移，小就左边右移
3. 如果相等，输出然后左边指针右移到第一个不是上个元素的位置，右边指针左移到第一个不是上个元素的位置
4. 如果左边指针大于右边指针则进行下一轮遍历

*/
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, n := range nums {
		// 如果是正数， 后面的无法是负数
		if nums[i] > 0 {
			break
		}
		// 如果这个数已经使用过，则跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - n
		l, r := i+1, len(nums)-1
		for l < r {
			if i == l {
				l++
			}
			if i == r {
				r--
			}
			tmp := nums[l] + nums[r]
			if tmp == target {
				// add
				result = append(result, []int{n, nums[l], nums[r]})
				// 如果左边这个数已经被使用过
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 如果右边这个数已经被使用过
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if tmp < target {
				l++
			} else {
				r--
			}

		}
	}
	return result
}

/*
删除排序数组中的重复项
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。

分析:
方法一：
	1. 定义两个指针，cur指向当前位置，ex指向下一个准备交换位置
    2. 如果ex == cur，ex就继续向前,不相等就cur前进一步，交换cur和ex，ex前进一步
	3. 最后返回cur+1
方法二：
	1. 定义快指针i和慢指针j，只要nums[i] == nums[j]，就增加j跳过重复项
    2. 如果不相等，就将i增加1，然后将nums[j]复制到nums[i]，然后继续遍历j
    3. 返回i + 1
*/

func removeDuplicates(nums []int) int {
	cur, ex := 0, 1
	for ex < len(nums) {
		if nums[ex] != nums[cur] {
			cur++
			nums[cur] = nums[ex]
			ex++
		} else {
			ex++
		}
	}
	return cur + 1
}

func removeDuplicates2(nums []int) int {
	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[i] == nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

/*
合并两个有序数组
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]

分析：
方法一：合并后排序  O(nlogn)
方法二：使用一个额外的数组，需要空间复杂度O(m)
方法三：从后往前遍历比较
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	// i for nums1, j for nums2
	i, j, pos := m-1, n-1, m+n-1
	for ; pos >= 0; pos-- {
		if i < 0 {
			nums1[pos] = nums2[j]
			j--
			continue
		}
		if j < 0 {
			nums1[pos] = nums1[i]
			i--
			continue
		}

		if nums1[i] >= nums2[j] {
			nums1[pos] = nums1[i]
			i--
		} else {
			nums1[pos] = nums2[j]
			j--
		}
	}
}
