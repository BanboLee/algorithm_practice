package string

/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例 1:
s = "leetcode"
返回 0.

示例 2:
s = "loveleetcode",
返回 2.

注意：您可以假定该字符串只包含小写字母。
分析：使用26位数组, 将出现的位置放置到数组中对应位置，第一个位置相等的就是所求。
*/

func firstUniqChar(s string) int {
	var lf [26]int
	for i, ch := range s {
		lf[ch-'a'] = i
	}
	for i, ch := range s {
		if i == lf[ch-'a'] {
			return i
		} else {
			lf[ch-'a'] = -1
		}
	}

	return -1
}
