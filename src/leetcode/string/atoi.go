package string

import (
	"math"
)

/*
字符串转整数

请你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，
作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231, 231 − 1]。如果数值超过这个范围，请返回 INT_MAX (231 − 1) 或 INT_MIN (−231) 。

示例 1:
输入: "42"
输出: 42

示例 2:
输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。

示例 3:
输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。

示例 4:
输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。

示例 5:
输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
     因此返回 INT_MIN 。

分析：
方法一：
1. 先决定符号，首个数字位置
2. 从前往后乘以10累加上去， 如果>max/10，这次运算肯定溢出
3. 直到首个非数字
方法二： 有限状态机
1. 程序对每次输入，分为4个状态： start、signed、in_number、end
2. 建立map表，对于每次输入决定下次正确输入，状态机如下表：
            ' '	    +/-	    number	    other
start       start	signed	in_number	end
signed	    end     end     in_number	end
in_number	end     end     in_number	end
end	        end     end     end         end

注意边界：数字大于INT_MAX或者小于INIT_MIN溢出的情况。
处理方法：提前预判，在就计算下一轮数字之前，将当前sum和INT_MIN/10以及INT_MAX/10比较。

*/

func myAtoi(str string) int {
	res := 0
	signed := 1
	i := 0

	// trim
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i >= len(str) {
		return 0
	}

	// signed
	if str[i] == '-' {
		signed = -1
	}
	if str[i] == '-' || str[i] == '+' {
		i++
	}

	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		r := str[i] - '0'
		println(r, res > math.MaxInt32/10, res == math.MaxInt32)
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && r > 7) {
			if signed == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(r)
		i++
	}

	if signed == 1 {
		return res
	}
	return -res
}

func atoiDFA(str string) int {
	var (
		start  = 0
		signed = 1
		inNum  = 2
		end    = 3
	)

	var table = map[int][]int{
		start:  {start, signed, inNum, end},
		signed: {end, end, inNum, end},
		inNum:  {end, end, inNum, end},
		end:    {end, end, end, end},
	}

	var res int
	si := 1
	state := start
	for _, ch := range str {
		var idx = 3
		if ch == ' ' {
			idx = 0
		} else if ch == '+' || ch == '-' {
			idx = 1
		} else if ch >= '0' && ch <= '9' {
			idx = 2
		}

		state = table[state][idx]
		switch state {
		case start:
			continue
		case signed:
			if ch == '-' {
				si = -1
			}
			break
		case inNum:
			r := ch - '0'
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && r > 7) {
				if si == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			res = res*10 + int(r)
			break
		case end:
			return si * res
		}
	}
	return si * res
}
