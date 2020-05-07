package string

/*
字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:
输入: num1 = "123", num2 = "456"
输出: "56088"
说明：
1.num1 和 num2 的长度小于110。
2.num1 和 num2 只包含数字 0-9。
3.num1 和 num2 均不以零开头，除非是数字 0 本身。
4.不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

分析：
1. 获得单个数字 ‘2’ - ‘0’
2. 完全手写竖式乘法 O(MN)
3. 优化竖式乘法 O(MN) 少做一次嵌套加法，实际时间会更小
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	// 数字2从后往前乘以第一个数字
	var carry int
	var prev = "0"
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := num2[i] - '0'
		l := len(num1) + 1 + carry
		var tmp = make([]byte, l, l)
		idx := l - 1 - carry
		for k := l - 1; k > idx; k-- {
			tmp[k] = '0'
		}
		var c byte = 0
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := num1[j] - '0'
			t1 := n1*n2 + c
			if t1 >= 10 {
				tmp[idx] = t1%10 + '0'
				c = t1 / 10
			} else {
				tmp[idx] = t1 + '0'
				c = 0
			}
			idx--
		}
		if c > 0 {
			tmp[0] = c + '0'
			prev = add(prev, string(tmp))
		} else {
			prev = add(prev, string(tmp[1:]))
		}
		carry++
	}
	return prev
}

// 按位加法
func add(s1, s2 string) string {
	num1 := []byte(s1)
	num2 := []byte(s2)
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	diff := len(num1) - len(num2)
	var carry byte = 0
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		var tmp byte
		if i-diff < 0 {
			tmp = n1 + carry
		} else {
			n2 := num2[i-diff] - '0'
			tmp = n1 + n2 + carry
		}
		if tmp >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		num1[i] = tmp%10 + '0'
	}

	if carry == 1 {
		return "1" + string(num1)
	} else {
		return string(num1)
	}
}

// 优化竖式乘法
func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var res = make([]byte, len(num1)+len(num2))
	for i := range res {
		res[i] = '0'
	}
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			tmp := res[i+j+1] - '0' + n1*n2 // 当前两数乘积+低位
			res[i+j+1] = tmp%10 + '0'       // 低位
			res[i+j] += tmp / 10            // 高位为原来值+进位

			// 细节：这里由于底层直接用了byte，为了后面直接转string，需要小心‘0’的使用。
		}
	}
	// 首位去0
	for i, v := range res {
		if v != '0' {
			return string(res[i:])
		}
	}
	return "0"
}

// 双指针加法 不转byte
func add1(s1, s2 string) string {
	lm := len(s1)
	if len(s2) > lm {
		lm = len(s2)
	}
	var res = make([]byte, lm+1) // 加法的结果顶多比原来最大的值多一位

	j, k := len(s1)-1, len(s2)-1
	var carry uint8 = 0
	for i := lm; i > 0; i-- {
		var n1 uint8 = 0
		var n2 uint8 = 0
		if j >= 0 {
			n1 = s1[j] - '0'
		}
		if k >= 0 {
			n2 = s2[k] - '0'
		}
		sum := n1 + n2 + carry
		res[i] = sum%10 + '0'
		carry = sum / 10
		j--
		k--
	}

	if carry > 0 {
		res[0] = '1'
		return string(res)
	} else {
		return string(res[1:])
	}
}
