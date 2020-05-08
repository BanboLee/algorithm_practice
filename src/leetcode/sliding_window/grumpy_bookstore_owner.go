package sliding_window

/*
leet-code 1052 爱生气的书店老板
链接：https://leetcode-cn.com/problems/grumpy-bookstore-owner

今天，书店老板有一家店打算试营业 customers.length 分钟。
每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，
否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

示例：
输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
输出：16
解释：
书店老板在最后 3 分钟保持冷静。
感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.

提示：
1 <= X <= customers.length == grumpy.length <= 20000
0 <= customers[i] <= 1000
0 <= grumpy[i] <= 1

分析:
滑动窗口
1. 对于正常的计算，每次滑动一个窗口，对出窗分钟，如果生气，选择不出窗，如果不生气，出窗相应客户

2. 找到最大的窗口。。。
*/

func maxSatisfied(customers []int, grumpy []int, X int) int {
	if len(customers) == 0 {
		return 0
	}

	var (
		windowSum int
		normalSum int
		prevMax   int
		prevSum   int
	)

	t := 0
	for ; t < X; t++ {
		windowSum += customers[t]
		if grumpy[t] == 0 {
			normalSum += customers[t]
		}
	}

	for {
		if windowSum+prevSum > prevMax {
			prevMax = windowSum + prevSum
		}

		if t >= len(customers) {
			break
		}

		// 出窗
		if grumpy[t-X] == 0 {
			prevSum += customers[t-X]
		}

		// 入窗
		windowSum -= customers[t-X]
		windowSum += customers[t]
		if grumpy[t] == 0 {
			prevMax += customers[t]
			normalSum += customers[t]
		}
		t++
	}

	return max(prevMax, normalSum)
}

// 还是想复杂了。
// 只需要补上窗口大小的顾客数就好了。
func maxSatisfied1(customers []int, grumpy []int, X int) int {
	var (
		sum, cur, maxCur int
	)

	for i := 0; i < len(customers); i++ {
		// 入窗
		if grumpy[i] == 0 {
			sum += customers[i]
		} else {
			cur += customers[i]
		}

		// 出窗
		if i >= X && grumpy[i-X] == 1 {
			cur -= customers[i-X]
		}

		if cur > maxCur {
			maxCur = cur
		}
	}
	return sum + maxCur
}
