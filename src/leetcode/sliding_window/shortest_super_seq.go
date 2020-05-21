package sliding_window

/*
leetcode 面试题 17.18. 最短超串
链接：https://leetcode-cn.com/problems/shortest-supersequence-lcci

假设你有两个数组，一个长一个短，短的元素均不相同。找到长数组中包含短数组所有的元素的最短子数组，其出现顺序无关紧要。
返回最短子数组的左端点和右端点，如有多个满足条件的子数组，返回左端点最小的一个。若不存在，返回空数组。

示例 1:
输入:
big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7]
small = [1,5,9]
输出: [7,10]

示例 2:
输入:
big = [1,2,3]
small = [4]
输出: []

提示：
big.length <= 100000
1 <= small.length <= 100000
*/

/*
   滑动窗口+hash
*/
func shortestSeq(big []int, small []int) []int {
	if len(small) > len(big) {
		return nil
	}

	mp := make(map[int]int)
	cnt := len(small)
	for _, v := range small {
		mp[v]++
	}

	l, r := 0, 0
	var res []int

	for ; r < len(big); r++ {
		// 入窗
		if mp[big[r]] > 0 {
			cnt--
		}
		mp[big[r]]--

		if cnt == 0 {
			if res == nil {
				res = append(res, l, r)
			} else {
				if r-l < res[1]-res[0] {
					res[0], res[1] = l, r
				}
			}
		}

		// 出窗
		for l <= r && (cnt == 0 || mp[big[l]] < 0) {
			if mp[big[l]] == 0 {
				cnt++
			}
			mp[big[l]]++
			l++
		}
	}
	return res
}
