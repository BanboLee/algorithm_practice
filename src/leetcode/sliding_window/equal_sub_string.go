package sliding_window

/*
leet-code 1208 尽可能使字符串相等
链接：https://leetcode-cn.com/problems/get-equal-substrings-within-budget/

给你两个长度相同的字符串，s 和 t。
将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。
如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

示例 1：
输入：s = "abcd", t = "bcdf", cost = 3
输出：3
解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。

示例 2：
输入：s = "abcd", t = "cdef", cost = 3
输出：1
解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。

示例 3：
输入：s = "abcd", t = "acde", cost = 0
输出：1
解释：你无法作出任何改动，所以最大长度为 1。

提示：
1 <= s.length, t.length <= 10^5
0 <= maxCost <= 10^6
s 和 t 都只含小写英文字母。

分析：
1. 双指针+滑动窗口
2. 生成差值数组，转换为最长连续子数组问题
*/
func equalSubstring(s string, t string, maxCost int) int {
	if s == t {
		return len(s)
	}

	maxLen := 0
	curCost, l := 0, 0
	for i := 0; i < len(s); i++ {
		var cost uint8
		if s[i] > t[i] {
			cost = s[i] - t[i]
		} else {
			cost = t[i] - s[i]
		}

		// 如果当前字符消耗超过所有
		if int(cost) > maxCost {
			if maxLen < i-l {
				maxLen = i - l
			}
			l = i + 1
			curCost = 0
			continue
		}

		// 调整窗口
		curCost += int(cost)
		if curCost > maxCost {
			if maxLen < i-l {
				maxLen = i - l
			}
			if s[l] > t[l] {
				curCost -= int(s[l] - t[l])
			} else {
				curCost -= int(t[l] - s[l])
			}
			l++
		} else if i == len(s)-1 {
			if maxLen < i-l+1 {
				maxLen = i - l + 1
			}
		}
	}

	return maxLen
}

// 使用转化
func equalSubstring1(s string, t string, maxCost int) int {
	var costs = make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > t[i] {
			costs[i] = s[i] - t[i]
		} else {
			costs[i] = t[i] - s[i]
		}
	}

	maxLen, curCost, l := 0, 0, 0
	for j, v := range costs {
		curCost += int(v)
		for curCost > maxCost {
			if maxLen < j-l {
				maxLen = j - l
			}
			curCost -= int(costs[l])
			l++
		}
	}

	if maxLen < len(s)-l {
		maxLen = len(s) - l
	}

	return maxLen
}
