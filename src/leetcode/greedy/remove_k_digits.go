package greedy

/*
leetcode 402. 移掉K位数字
链接：https://leetcode-cn.com/problems/remove-k-digits

给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:
num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。

示例 1 :
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/

// 贪心算法，每次移除一个以后后面的变成新问题求最优解
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	var stack []byte
	cnt := k
	for i := range num {
		digit := num[i]
		// 弹出小于k个的之前元素，由于是按栈排列，所剩的必然是之前最小的数字
		for len(stack) != 0 && cnt > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			cnt--
		}
		stack = append(stack, digit)
	}

	// 去掉末尾被加上的数字
	stack = stack[:len(num)-k]
	var res []byte
	flag := true
	for _, ch := range stack {
		if ch == '0' && flag {
			continue
		}
		flag = false
		res = append(res, ch)
	}

	if len(res) == 0 {
		return "0"
	}
	return string(res)

}
