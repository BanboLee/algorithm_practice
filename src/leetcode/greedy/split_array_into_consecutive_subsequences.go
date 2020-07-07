package greedy

/*
leetcode 659. 分割数组为连续子序列
链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences

给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。
如果可以完成上述分割，则返回 true ；否则，返回 false 。

示例 1：
输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3
3, 4, 5

示例 2：
输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3, 4, 5
3, 4, 5

示例 3：
输入: [1,2,3,4,4,5]
输出: False

提示：
输入的数组长度范围为 [1, 10000]
*/

func isPossible(nums []int) bool {
	count, tails := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	for _, num := range nums {
		if count[num] == 0 {
			// 没有可用的该数字
			continue
		} else if tails[num] > 0 {
			// 有num-1为尾巴的链
			tails[num+1]++
			tails[num]--
		} else if count[num+1] > 0 && count[num+2] > 0 {
			// 新建一条链
			tails[num+3]++
			count[num+1]--
			count[num+2]--
		} else {
			// 无法加入一条有的链并无法创建一条新的链
			return false
		}
		count[num]--
	}

	return true
}
