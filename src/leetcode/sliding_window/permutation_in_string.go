package sliding_window

/*
leet-code 567 字符串的排列
链接：https://leetcode-cn.com/problems/permutation-in-string

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False

注意：
输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
*/

func checkInclusion(s1 string, s2 string) bool {
	// 使用数组进行所需字符计数
	var table [26]int
	for _, ch := range s1 {
		table[ch-'a']++
	}

	l, r := 0, 0
	for ; r < len(s2); r++ {
		// 入窗
		table[s2[r]-'a']--

		// 如若当前字符数超过所需， 则需要出窗左侧直至第一个该字符
		for table[s2[r]-'a'] < 0 && l <= r {
			table[s2[l]-'a']++
			l++
		}

		// 如果窗口正好匹配大小，那么得到结果
		if r-l+1 == len(s1) {
			return true
		}
	}
	return false
}
