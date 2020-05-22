package prefix_and

/*
leetcode 307. 区域和检索 - 数组可修改
链接：https://leetcode-cn.com/problems/range-sum-query-mutable

给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。

示例:
Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8

说明:
数组仅可以在 update 函数下进行修改。
你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。

1. O(n)  更新数组, 迭代求和
2. TODO: 线段树

*/

type NumArray1 struct {
	data []int
}

func Constructor2(nums []int) NumArray1 {
	na := NumArray1{nums}
	return na
}

func (this *NumArray1) Update(i int, val int) {
	this.data[i] = val
}

func (this *NumArray1) SumRange1(i int, j int) int {
	var res int
	for _, v := range this.data {
		res += v
	}
	return res
}
