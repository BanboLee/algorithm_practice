package sliding_window

/*
leet-code 395 至少有K个重复字符的最长子串
链接：https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/

找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。

示例 1:
输入:
s = "aaabb", k = 3
输出:
3
最长子串为 "aaa" ，其中 'a' 重复了 3 次。

示例 2:
输入:
s = "ababbc", k = 2
输出:
5
最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

分析：分治法 O(nlogn)
1. 出现不符合条件的字符时，可行条件必在两边，分为两个字符串处理
2. 分割时，必要找到下一个可行字符，减少多次运算
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	var table [26]int
	for _, ch := range s {
		table[ch-'a']++
	}

	maxLen, l := 0, 0
	// 用于标识第一个出现的不符合条件的字符
	c := -1
	// 用于标识是否出现分割，如果没有出现，那么长度必为字符串长度
	split := false
	for i := 0; i < len(s); i++ {
		// 第一次出现不符合的字符，标记一下位置
		if table[s[i]-'a'] < k && c == -1 {
			c = i
			continue
		}
		// 恢复可行字符时，开始做分割，此时从c到当前位置的字符串是不可用的，不需要理会
		if table[s[i]-'a'] >= k && c != -1 {
			maxLen = max(maxLen, max(longestSubstring(s[l:c], k), longestSubstring(s[i:], k)))
			// 下一次分割必然是从该字符往右的连续字符，毕竟之前的都不可用
			l = i + 1
			// 标记已进行分割
			split = true
			// 恢复异常字符标识位
			c = -1
		}
	}

	// 处理最后一次分割
	if c != -1 {
		split = true
		maxLen = max(maxLen, longestSubstring(s[l:c], k))
	}

	// 处理未分割情况
	if !split {
		maxLen = len(s)
	}

	return maxLen
}
