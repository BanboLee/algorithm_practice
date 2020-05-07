package array

import (
	"sort"
)

/*
杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
分析：
纯粹找规律撸代码
*/
func generate(numRows int) [][]int {
	var result = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		l := i + 1
		var cur = make([]int, l)
		cur[0] = 1
		cur[l-1] = 1
		for j := 1; j < l-1; j++ {
			cur[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = cur
	}

	return result
}

/*
缺失数字
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:
输入: [3,0,1]
输出: 2

示例 2:
输入: [9,6,4,2,3,5,7,0,1]
输出: 8

说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

分析:
1. 所有数字与数组中数字做异或运算
2. 等差数列求和减去数组中每个数字

*/

func missingNumber(nums []int) int {
	// bit operation.
	var n = 0
	// for i, v := range nums {
	// 	n ^= (i + 1) ^ v
	// }

	// Arithmetic sequence。
	for i, v := range nums {
		n += (i + 1) - v

	}

	return n
}

/*
旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

说明:
尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的原地算法。

分析：
1.使用环装替换 O(n) O(1)
2.使用翻转 O(n) O(1)
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	// 替换次数为数组长度
	count := 0
	for start := 0; count < len(nums); start++ {
		current := (start + k) % len(nums)
		prev := nums[start]
		for start != current {
			tmp := nums[current]
			nums[current] = prev
			prev = tmp
			count++
			current = (current + k) % len(nums)
		}
		nums[current] = prev
		count++
	}
}

func rotate2(nums []int, k int) {
	reverse(nums)
	m := k % len(nums)
	reverse(nums[:m])
	reverse(nums[m:])
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
合并区间
给出一个区间的集合，请合并所有重叠的区间。

示例 1:
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

分析：
1. 使用图算法中的邻接表作连通块
2. 先把原有数组按第一个数排序，然后合并
*/
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var (
		result [][]int
		idx    = 0
	)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if result[idx][1] >= intervals[i][1] {
			continue
		} else if intervals[i][0] <= result[idx][1] {
			result[idx][1] = intervals[i][1]
		} else {
			result = append(result, intervals[i])
			idx++
		}
	}
	return result
}

/*
下一个排列
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2 3,2,1 → 1,2,3 1,1,5 → 1,5,1

分析：
多举例子找规律
再举几个例子验证仍然符合这个规律。因此，处理过程可总结以下几个步骤：
1.从后往前遍历数组，找到数值开始减小的那个位置idx，即nums[idx-1] < nums[idx]；
2.令idx_l = idx-1，从idx位置从前往后移动直到找到最后一个数值大于nums[idx_l]的元素；
3.交换nums[idx_l]和nums[idx]的值；
4.反转数组在idx位置之后的所有元素。
*/
func nextPermutation(nums []int) {
	// 从右向左找到第一个变小的位置
	var l = -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			l = i - 1
			break
		}
	}
	if l == -1 {
		reserve(nums)
		return
	}

	// 从该位置开始，从左向右找到最后一个它大的值
	var r = -1
	for j := l + 1; j < len(nums); j++ {
		if nums[j] > nums[l] {
			r = j
		}
	}
	println(l, r)
	// 交换并翻转从该位置往后的元素
	nums[l], nums[r] = nums[r], nums[l]
	reserve(nums[l+1:])
}

func reserve(data []int) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

/*
寻找重复数
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:
输入: [1,3,4,2,2]
输出: 2

示例 2:
输入: [3,1,3,4,2]
输出: 3

说明：
不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。

分析:
1. 环状替换
2. floyd算法，检测环

*/

func findDuplicate(nums []int) int {
	for i, v := range nums {
		if v == i+1 {
			continue
		}

		for nums[i] != i+1 {
			if nums[nums[i]-1] == nums[i] {
				return nums[i]
			} else {
				nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			}
		}
	}
	return 0
}

/*
floyd的龟兔赛跑法
假设环入口为m，环周长为c
快指针每次前进2个位置，慢指针每次前进一个位置，那么最终相遇时，慢指针前进n步，快指针则前进2n步
快指针比慢指针多n步，而这n步说明快指针一直在环内，所以n%c==0
慢指针在环里前进的步数为n-m，现在让find指针前进m步，慢指针则n-m+m到达环的入口处，此时即找到环的入口
*/
func findDuplicate2(nums []int) int {
	fast := nums[nums[nums[0]]]
	low := nums[nums[0]]
	for fast != low {
		fast = nums[nums[fast]]
		low = nums[low]
	}
	ptr1, ptr2 := nums[0], low
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr2
}

/*
缺失的第一个正数 (hard)
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:
输入: [1,2,0]
输出: 3

示例 2:
输入: [3,4,-1,1]
输出: 2

示例 3:
输入: [7,8,9,11,12]
输出: 1

说明:
你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

分析:
1. 不可能为负数、0和大于n+1的数字，可以先排除，全部设置为1
2. 对每个元素v，将下标为v-1的位置置为负表示其出现过(hash)
*/
func firstMissingPositive(nums []int) int {
	has1 := false
	for _, v := range nums {
		if v == 1 {
			has1 = true
			break
		}
	}

	if !has1 {
		return 1
	}
	l := len(nums)
	for i, v := range nums {
		if v <= 0 || v >= l+1 {
			nums[i] = 1
		}
	}
	for _, v := range nums {
		if v < 0 {
			v = 0 - v
		}
		if nums[v-1] < 0 {
			continue
		}
		nums[v-1] = 0 - nums[v-1]
	}

	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return l + 1
}
