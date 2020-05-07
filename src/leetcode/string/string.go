package string

/*
机器能否返回原点
在二维平面上，有一个机器人从原点 (0, 0) 开始。给出它的移动顺序，判断这个机器人在完成移动后是否在 (0, 0) 处结束。
移动顺序由字符串表示。字符 move[i] 表示其第 i 次移动。机器人的有效动作有 R（右），L（左），U（上）和 D（下）。
如果机器人在完成所有动作后返回原点，则返回 true。否则，返回 false。
注意：机器人“面朝”的方向无关紧要。 “R” 将始终使机器人向右移动一次，“L” 将始终向左移动等。此外，假设每次移动机器人的移动幅度相同。

示例 1:
输入: "UD"
输出: true
解释：机器人向上移动一次，然后向下移动一次。所有动作都具有相同的幅度，因此它最终回到它开始的原点。因此，我们返回 true。
示例 2:
输入: "LL"
输出: false
解释：机器人向左移动两次。它最终位于原点的左侧，距原点有两次 “移动” 的距离。我们返回 false，因为它在移动结束时没有返回原点。
*/

func judgeCircle(moves string) bool {
	u, r := 0, 0
	for _, s := range moves {
		switch s {
		case 'U':
			u++
		case 'D':
			u--
		case 'L':
			r--
		case 'R':
			r++
		}
	}

	if u == 0 && r == 0 {
		return true
	}
	return false
}

/*
反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]

*/

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

/*
反转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。

示例:
输入: "the sky is blue",
输出: "blue is sky the".
说明:
无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
进阶: 请选用C语言的用户尝试使用 O(1) 空间复杂度的原地解法。
*/
func reverseWords(s string) string {
	var result []byte
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if cnt == 0 {
				continue
			}
			if len(result) != 0 {
				result = append(result, ' ')
			}
			tmp := 1
			for cnt > 0 {
				result = append(result, s[i+tmp])
				cnt--
				tmp++
			}
		} else {
			cnt++
		}
	}

	if cnt > 0 {
		if len(result) != 0 {
			result = append(result, ' ')
		}
		for i := 0; i < cnt; i++ {
			result = append(result, s[i])
		}
	}

	return string(result)
}

/*
反转字符串中的单词III

给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:
输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*/
func reverseWords3(s string) string {
	l, r := 0, 0
	var res = make([]byte, len(s), len(s))
	for r <= len(s) {
		if r == len(s) || s[r] == ' ' {
			j := 1
			for l != r {
				res[l] = s[r-j]
				l++
				j++
			}
			if r != len(s) {
				res[r] = ' '
			}
			r++
			l = r
		} else {
			r++
		}
	}
	return string(res)
}

/*
罗马数字转整数

罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做 XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，
所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1：
输入: "IX"
输出: 9

示例 2:
输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.

示例 3:
输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

分析：
1. 暴力法
2. 构建完全字典法
3. 加减法

规律，大值左边出现小值做减法。
*/

var romanInteger = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	i := 0
	r := 0
	for i < len(s) {
		if s[i] == 'I' && i != len(s)-1 {
			if s[i+1] == 'X' {
				r += 9
				i += 2
			} else if s[i+1] == 'V' {
				r += 4
				i += 2
			} else {
				r += 1
				i++
			}
		} else if s[i] == 'X' && i != len(s)-1 {
			if s[i+1] == 'L' {
				r += 40
				i += 2
			} else if s[i+1] == 'C' {
				r += 90
				i += 2
			} else {
				r += 10
				i++
			}
		} else if s[i] == 'C' && i != len(s)-1 {
			if s[i+1] == 'D' {
				r += 400
				i += 2
			} else if s[i+1] == 'M' {
				r += 900
				i += 2
			} else {
				r += 100
				i++
			}
		} else {
			r += romanInteger[s[i]]
			i++
		}
	}
	return r
}
