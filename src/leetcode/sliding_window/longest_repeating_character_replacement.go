package sliding_window

/*
leet-code 424 替换后的最长重复字符
链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement

给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，
总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
注意:
字符串长度 和 k 不会超过 104。

示例 1:
输入:
s = "ABAB", k = 2
输出:
4
解释:
用两个'A'替换为两个'B',反之亦然。

示例 2:
输入:
s = "AABABBA", k = 1
输出:
4
解释:
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

分析:
*/

/*
  总结：
    使用字符集和字符次数来做窗口是最合适的！！！
    本题的重点就在于每一个窗口中，需要替换次数最少就能达到最大长度的，肯定是窗口中出现次数最多的。
*/
func characterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}

	// 使用字符集当窗口，记录本窗口中各个字符出现的次数
	var table [26]int
	// maxCnt 记录当前窗口中出现最大的次数
	maxLen, l, r, maxCnt := 0, 0, 0, 0
	for ; r < len(s); r++ {
		// 入窗
		table[s[r]-'A']++
		// 如果窗口中最大次数发生变化，改变最大值
		if maxCnt < table[s[r]-'A'] {
			maxCnt = table[s[r]-'A']
		}

		// 如果替换次数超过了k，就将最左边的字符出窗
		if r-l+1-maxCnt > k {
			table[s[l]-'A']--
			l++
		}

		// 这个时候虽然不一定最左边的是需要占用替换名额的，但其最终长度依然与上一个有效长度一致，
		// 不会发生变化，直到窗口最大次数字符发生变化
		if maxLen < r-l+1 {
			maxLen = r - l + 1
		}
	}

	return maxLen
}
