package string

import "fmt"

/*
最小覆盖子串
给定一个字符串 S 和一个字符串 T，请在 S 中找出包含 T 所有字母的最小子串。

示例：
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：
如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

分析：
1. 使用双指针伸开寻找窗口
2. 滑动窗口, 入窗和出窗
3. 收缩窗口
*/

func minWindow(s string, t string) string {
	var table = make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		table[t[i]]++
	}
	cnt := len(t)

	l, r := 0, 0
	var res = ""
	for r < len(s) {
		vr, ok := table[s[r]]
		if ok {
			if vr > 0 {
				cnt--
			}
			table[s[r]]--
		}
		r++
		for cnt == 0 {
			v, ok := table[s[l]]
			if ok {
				if v >= 0 {
					cnt++
					table[s[l]]++
					if len(res) == 0 || len(s[l:r]) < len(res) {
						res = s[l:r]
					}
				} else {
					table[s[l]]++
				}
			}
			l++
		}
		fmt.Printf("l=%d, r=%d, cnt=%d, sr=%c, table=%v\n", l, r, cnt, s[r-1], table)
	}
	return res
}

func minWindow1(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	l, count, max, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		// 1. 入窗： 若某个位置依然需要的数目大于等于0， cnt--
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		// 2. 出窗： 左边界小于右边界的情况下，如果某个位置小于0，说明这个位置的后面依然充足，
		// 可以出窗来缩小窗口
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		// 3. 如果count=0,说明当前窗口正好符合所求
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}

	return results
}
