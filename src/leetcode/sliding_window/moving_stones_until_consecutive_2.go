package sliding_window

import (
	"fmt"
	"sort"
)

/*
leetcode 1040. 移动石子直到连续 II
链接：https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii

在一个长度无限的数轴上，第 i 颗石子的位置为 stones[i]。如果一颗石子的位置最小/最大，那么该石子被称作端点石子。
每个回合，你可以将一颗端点石子拿起并移动到一个未占用的位置，使得该石子不再是一颗端点石子。
值得注意的是，如果石子像 stones = [1,2,5] 这样，你将无法移动位于位置 5 的端点石子，
因为无论将它移动到任何位置（例如 0 或 3），该石子都仍然会是端点石子。当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。
要使游戏结束，你可以执行的最小和最大移动次数分别是多少？ 以长度为 2 的数组形式返回答案：answer = [minimum_moves, maximum_moves] 。

示例 1：
输入：[7,4,9]
输出：[1,2]
解释：
我们可以移动一次，4 -> 8，游戏结束。
或者，我们可以移动两次 9 -> 5，4 -> 6，游戏结束。

示例 2：
输入：[6,5,4,3,10]
输出：[2,3]
解释：
我们可以移动 3 -> 8，接着是 10 -> 7，游戏结束。
或者，我们可以移动 3 -> 7, 4 -> 8, 5 -> 9，游戏结束。
注意，我们无法进行 10 -> 2 这样的移动来结束游戏，因为这是不合要求的移动。

示例 3：
输入：[100,101,104,102,103]
输出：[0,0]

提示：
3 <= stones.length <= 10^4
1 <= stones[i] <= 10^9
stones[i] 的值各不相同。
*/

/*
   这题就难在读懂题目上，题目的意思是，把最小或者最大的插到中间，让它不再在两端。
   最大值：
        1. 已占用位置为N, 目前总共占用的长度为stones[N-1] - stones[0] + 1,
           减去已占用位置剩下的为：
           s1 = stones[N-1] - stones[0] + 1 - N
        2. 由于第一个和第二个石子之间不能放置第一个石子，同理最后一个也不行，
           那么这中间需要去掉的最少空间为:
           s2 = min(stones[N-1] - stones[N - 2] - 1, stones[1]-stones[0] - 1)
        3. 这样这个最大值就是 s1 - s2
    最小值：
        我们只需要找到一个连续的可以装下所有石子的位置，相当于滑动窗口寻找满足要求的子数组。

*/
func numMovesStonesII(stones []int) []int {
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] < stones[j]
	})

	fmt.Printf("after sort: %v\n", stones)
	n := len(stones)
	// 最大值
	maxEmpty := stones[n-1] - stones[0] + 1 - n
	maxEmpty -= min(stones[n-1]-stones[n-2]-1, stones[1]-stones[0]-1)
	var res = make([]int, 2)
	res[0] = maxEmpty
	res[1] = maxEmpty

	l, r := 0, 0
	for ; r < n; r++ {
		// 出窗, 我们需要的是小于等于n的情况
		for stones[r]-stones[l]+1 > n {
			l++
		}

		// 特殊情况， 前面全有序
		if n-(r-l+1) == 1 && stones[r]-stones[l]+1 == n-1 {
			res[0] = min(res[0], 2)
		} else {
			res[0] = min(res[0], n-(r-l+1))
		}
	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
