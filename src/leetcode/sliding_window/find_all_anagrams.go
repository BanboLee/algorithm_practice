package sliding_window

/*
leet-code 438 找到字符串中所有字母异位词
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

示例 1:
输入:
s: "cbaebabacd" p: "abc"
输出:
[0, 6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

示例 2:
输入:
s: "abab" p: "ab"
输出:
[0, 1, 2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

分析：
字符计数来做滑动窗口
*/

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(s) < len(p) {
		return nil
	}

	// 1.使用map存储字符次数
	var table = make(map[uint8]int)
	for i := range p {
		table[p[i]-'a']++
	}
	var res []int
	l, cnt := 0, len(p)
	for i := range s {
		v, ok := table[s[i]-'a']
		// 若是其他字符，返还前面连续的字符次数
		if !ok {
			for l < i {
				table[s[l]-'a']++
				cnt++
				l++
			}
			l = i + 1
			continue
		}

		// 无法再入窗时， 把左边直到第一个该字符的所有字符出窗
		if v == 0 {
			for s[l] != s[i] {
				table[s[l]-'a']++
				cnt++
				l++
			}
			table[s[l]-'a']++
			cnt++
			l++
		}

		// 入窗
		table[s[i]-'a']--
		cnt--

		// 出窗
		if cnt == 0 {
			res = append(res, l)
			table[s[l]-'a']++
			l++
			cnt++
		}
	}

	return res
}

func findAnagrams1(s string, p string) []int {
	if len(s) == 0 || len(s) < len(p) {
		return nil
	}

	var table [26]int
	for _, ch := range p {
		table[ch-'a']++
	}

	var res []int
	l, r := 0, 0
	for ; r < len(s); r++ {
		table[s[r]-'a']--

		for l <= r && table[s[r]-'a'] < 0 {
			table[s[l]-'a']++
			l++
		}

		if r-l+1 == len(p) {
			res = append(res, l)
		}
	}

	return res
}
