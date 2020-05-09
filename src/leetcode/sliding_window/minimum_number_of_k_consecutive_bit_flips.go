package sliding_window

/*
leetcode 995. K 连续位的最小翻转次数
链接：https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips

在仅包含 0 和 1 的数组 A 中，一次 K 位翻转包括选择一个长度为 K 的（连续）子数组，
同时将子数组中的每个 0 更改为 1，而每个 1 更改为 0。
返回所需的 K 位翻转的次数，以便数组没有值为 0 的元素。如果不可能，返回 -1。

示例 1：
输入：A = [0,1,0], K = 1
输出：2
解释：先翻转 A[0]，然后翻转 A[2]。

示例 2：
输入：A = [1,1,0], K = 2
输出：-1
解释：无论我们怎样翻转大小为 2 的子数组，我们都不能使数组变为 [1,1,1]。

示例 3：
输入：A = [0,0,0,1,0,1,1,0], K = 3
输出：3
解释：
翻转 A[0],A[1],A[2]: A变成 [1,1,1,1,0,1,1,0]
翻转 A[4],A[5],A[6]: A变成 [1,1,1,1,1,0,0,0]
翻转 A[5],A[6],A[7]: A变成 [1,1,1,1,1,1,1,1]

提示：
1 <= A.length <= 30000
1 <= K <= A.length

*/

/*
   滑动窗口解法
   第i个位置是否需要翻转取决于前面k-1个位置是否进行了翻转,
   如果前k-1个位置已经帮它进行了必要的翻转，则不需要再进行了，直接下一个，
   如果没有帮它进行翻转，则他自己翻转的同时，也要记着，i + k - 1位置也被翻转了。
*/
func minKBitFlips(A []int, K int) int {
	var (
		res    int
		flip   int                   // 记录当前窗口目前总共进行了多少次翻转
		window = make([]int, len(A)) // 记录某个位置是否进行了翻转
	)

	for i := 0; i < len(A); i++ {
		// 出窗
		if i >= K {
			flip -= window[i-K]
		}

		// 入窗
		if (flip%2)^A[i] == 0 {
			// 如果过了i+k-1还要修改后面窗口，那就不行
			if i+K > len(A) {
				return -1
			}
			flip++
			res++
			window[i] = 1
		}
	}

	return res
}