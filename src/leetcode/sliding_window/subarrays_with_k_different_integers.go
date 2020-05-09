package sliding_window

/*
leetcode 992. K 个不同整数的子数组
链接：https://leetcode-cn.com/problems/subarrays-with-k-different-integers

给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。
（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
返回 A 中好子数组的数目。

示例 1：
输入：A = [1,2,1,2,3], K = 2
输出：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
示例 2：
输入：A = [1,2,1,3,4], K = 3
输出：3
解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].

提示：
1 <= A.length <= 20000
1 <= A[i] <= A.length
1 <= K <= A.length

*/

/*
	这种解法，使用hash得到有效区间(i, j)内所有成立的子数组，
    外层出入窗固定时间复杂度为O(N), 对于寻找所有(i, j)内的有效子数组，
	最坏情况下，左边指针一直在最左边，没有出窗，时间复杂度为O(N²)
*/
func subarraysWithKDistinct(A []int, K int) int {
	var res, l int
	var table = make(map[int]int)
	for i := 0; i < len(A); i++ {
		// 入窗, 每个数字入窗一次  O(N)
		table[A[i]]++

		// 出窗, O(N)
		for len(table) > K {
			table[A[l]]--
			if table[A[l]] == 0 {
				delete(table, A[l])
			}
			l++
		}

		tmp := l
		// 使用双指针+滑动窗口确定所有可能, 最坏情况O(n²）
		for tmp <= i && len(table) == K {
			res++
			table[A[tmp]]--
			if table[A[tmp]] == 0 {
				delete(table, A[tmp])
			}
			tmp++
		}

		// 恢复hashtable
		tmp--
		for tmp >= l {
			table[A[tmp]]++
			tmp--
		}
	}
	return res
}

/*
	官方题解：
	定义前后两个窗口,后窗口left1和前窗口left2，后窗口合法的情况下，前窗口必合法,
    固定后窗口为合法子数组，前窗口为临界窗口，
    前窗口移动一次，必然要恢复到不合法的临界位置，left2-left1便是临界之前所有合法子数组数。
	这题关键要找到合法子数组之间的关系。
*/
type window map[int]int

func (w window) add(v int) {
	w[v]++
}

func (w window) remove(v int) {
	w[v]--
	if w[v] == 0 {
		delete(w, v)
	}
}

func subarraysWithKDistinct1(A []int, K int) int {
	var (
		w1    = make(window)
		w2    = make(window)
		left1 = 0
		left2 = 0
		res   = 0
	)

	for i := 0; i < len(A); i++ {
		w1.add(A[i])
		w2.add(A[i])

		// 后窗口[left1, A[i]]一定一直是合法的
		for len(w1) > K {
			w1.remove(A[left1])
			left1++
		}

		// 前窗口保持临界状态，也即得到当前临界窗口下，
		// [left1, left2)所有到A[i]的数组一定是合法的，
		for len(w2) >= K {
			w2.remove(A[left2])
			left2++
		}

		// [left1, left2]为之间所有合法的解
		res += left2 - left1
	}

	return res
}
