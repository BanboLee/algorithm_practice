package greedy

import "sort"

/*
leetcode 435. 无重叠区间
链接：https://leetcode-cn.com/problems/non-overlapping-intervals

给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:
可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
示例 1:
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。
示例 2:
输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
示例 3:
输入: [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
*/

/*
	如果后区间的终点在当前终点之前，直接去掉前区间，会有更大空间容纳更多区间
	如果后区间终点和起点在当前区间终点之后，无法去掉空间
	如果后区间的起点在当前区间终点之前并且终点在当前终点之后，使用贪心，直接去掉后区间
*/
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 按起点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := 0
	idx := 0

	for i := 1; i < len(intervals); i++ {
		// 如果下一区间的起点在当前区间之前
		if intervals[i][0] < intervals[idx][1] {
			// 两种情况
			if intervals[i][1] < intervals[idx][1] {
				// 如果当前区间的终点在下一区间终点之后， 去掉当前区间
				idx = i
			}
			// 如果当前区间的终点在下一区间终点之前，去掉下一区间
			ans++
		} else {
			// 这时两个区间不重叠，无法去掉任何一个
			idx = i
		}
	}

	return ans
}
