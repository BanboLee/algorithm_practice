package segment_tree

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
*/

type NumArray struct {
	preSum []int
	data   []int
}

// 构建线段树
func (this *NumArray) build(begin, end, root int) {
	if begin == end {
		this.preSum[root] = this.data[begin]
		return
	}

	mid := (end + begin) / 2
	this.build(begin, mid, root*2+1)
	this.build(mid+1, end, root*2+2)

	this.preSum[root] = this.preSum[root*2+1] + this.preSum[root*2+2]
}

// 查询区间[l, r]的和
func (this *NumArray) getSum(l, r, s, t, p int) int {
	if l <= s && t <= r {
		return this.preSum[p]
	}

	mid := (s + t) / 2
	sum := 0
	if l <= mid {
		sum += this.getSum(l, r, s, mid, p*2+1)
	}
	if r > mid {
		sum += this.getSum(l, r, mid+1, t, p*2+2)
	}
	return sum
}

func Constructor(nums []int) NumArray {
	na := NumArray{make([]int, len(nums)*4), nums}
	if len(nums) != 0 {
		na.build(0, len(na.data)-1, 0)
	}
	return na
}

func (this *NumArray) Update(i int, val int) {
	// 单个的某个元素修改，需要其父亲都被修改，重建
	if i < 0 || i >= len(this.data) {
		return
	}
	this.data[i] = val
	this.build(0, len(this.data)-1, 0)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.getSum(i, j, 0, len(this.data)-1, 0)
}
