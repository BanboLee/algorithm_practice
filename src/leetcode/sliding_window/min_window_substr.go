package sliding_window

/*
leet-code 76 最小覆盖子串
链接：https://leetcode-cn.com/problems/minimum-window-substring

给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：
如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

分析：
滑动窗口
最基本套路，先把入窗和出窗操作做了，然后再讨论结果
*/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	var table [128]int
	for i := range t {
		table[t[i]]++
	}
	var res = ""
	cnt, max := len(t), len(s)+1
	l, r := 0, 0
	for ; r < len(s); r++ {
		// 入窗， 有效字符则cnt--
		table[s[r]]--
		if table[s[r]] >= 0 {
			cnt--
		}

		// 出窗， 若左边界次数小于0，说明窗中该字符充足，进行出窗操作即可
		for l < r && table[s[l]] < 0 {
			table[s[l]]++
			l++
		}

		// 如果当前窗口正好符合， 则更新结果
		if cnt == 0 && max > r-l+1 {
			max = r - l + 1
			res = s[l : r+1]
		}
	}

	return res
}
