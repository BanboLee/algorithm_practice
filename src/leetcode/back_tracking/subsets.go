package back_tracking

/*
leetcode 78. 子集
链接：https://leetcode-cn.com/problems/subsets

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。

示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsetsRecursion(nums []int, res *[][]int, cur []int, index int) {
	if index == len(nums) {
		*res = append(*res, cur)
		return
	}

	subsetsRecursion(nums, res, cur, index+1)

	cur = append(cur, nums[index])
	subsetsRecursion(nums, res, cur, index+1)
	cur = cur[:len(cur)-1]
}

func subsets(nums []int) [][]int {
	var ans [][]int
	var tmp []int

	subsetsRecursion(nums, &ans, tmp, 0)
	return ans
}

// 每次往结果中的每一个子集添加新的数
func subsets1(nums []int) [][]int {
	res := [][]int{{}}
	for _, v := range nums {
		for _, cur := range res {
			tmp := make([]int, len(cur)+1)
			copy(tmp, cur)
			tmp[len(tmp)-1] = v
			res = append(res, tmp)
		}
		res = append(res, []int{v})
	}
	return res
}
