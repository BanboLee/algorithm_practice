package sliding_window

/*
leetcode 1371. 每个元音包含偶数次的最长子字符串
链接：https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts

给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：
每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。

示例 1：
输入：s = "eleetminicoworoep"
输出：13
解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。

示例 2：
输入：s = "leetcodeisgreat"
输出：5
解释：最长子字符串是 "leetc" ，其中包含 2 个 e 。

示例 3：
输入：s = "bcbcbc"
输出：6
解释：这个示例中，字符串 "bcbcbc" 本身就是最长的，因为所有的元音 a，e，i，o，u 都出现了 0 次。

提示：
1 <= s.length <= 5 x 10^5
s 只包含小写英文字母。
*/

/*
   入窗： 如果是元音字母则mp[ch]++, 如果是偶数，flag--,奇数flag--
   flag=0: 如果当前窗口比上个窗口大，更新结果
   超时  O(n²)

*/
func findTheLongestSubstring(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		flag := 0
		var mp [26]int
		mp['a'-'a'] = 1
		mp['e'-'a'] = 1
		mp['i'-'a'] = 1
		mp['o'-'a'] = 1
		mp['u'-'a'] = 1

		for r := i; r < len(s); r++ {
			// 入窗
			if mp[s[r]-'a'] > 0 {
				mp[s[r]-'a']++
				if mp[s[r]-'a']%2 == 0 {
					flag++
				} else {
					flag--
				}
			}

			if flag == 0 && res < r-i+1 {
				res = r - i + 1
			}
		}
	}

	return res
}

/*
	官方解答： 前缀和+状态压缩
	前缀和：根据之前的每个字母的状态，更新当前位置是奇数还是偶数,
			得到一个状态，如果该状态出现过，那么从第一次出现的位置到当前是符合要求的
			偶-偶=偶， 奇-奇=偶
	状态压缩： 将5个字母压缩到5个比特位上
	O(n)
*/
func findTheLongestSubstring1(s string) int {
	var (
		res    int
		status uint8
		pos    [32]int
	)

	for i := range pos {
		pos[i] = -1
	}
	pos[0] = 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}

		if pos[status] != -1 {
			// 如果状态已经出现过，则从第一个位置到当前位置的子串，必然符合要求
			res = max(res, i+1-pos[status])
		} else {
			// 状态为status时有i+1个字符， 记录某个状态第一次出现的位置
			pos[status] = i + 1
		}
	}
	return res
}
