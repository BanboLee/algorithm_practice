package string

/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:
所有输入只包含小写字母 a-z 。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prev := []byte(strs[0])
	n := 1
	for n < len(strs) {
		var l int
		if len(prev) < len(strs[n]) {
			l = len(prev)
		} else {
			l = len(strs[n])
		}

		i := 0
		for ; i < l; i++ {
			if prev[i] != strs[n][i] {
				prev = prev[:i]
				break
			}
		}
		if i == len(strs[n]) {
			prev = []byte(strs[n])
		}

		if len(prev) == 0 {
			return ""
		}
		n++
	}

	return string(prev)
}
