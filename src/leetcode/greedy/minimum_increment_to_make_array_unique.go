package greedy

import "sort"

/*
leetcode  945. 使数组唯一的最小增量
链接：https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique

给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。
返回使 A 中的每个值都是唯一的最少操作次数。

示例 1:
输入：[1,2,2]
输出：1
解释：经过一次 move 操作，数组将变为 [1, 2, 3]。

示例 2:
输入：[3,2,1,2,1,7]
输出：6
解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。

提示：
0 <= A.length <= 40000
0 <= A[i] < 40000
*/

// 计数暴力加，超时
func minIncrementForUnique(A []int) int {
	var mp = make(map[int]int)

	ans := 0

	for _, v := range A {
		for mp[v] > 0 {
			v++
			ans++
		}
		mp[v] = 1
	}

	return ans
}

/*
	计数法优化：
		先计数，如果某个数第一次出现，就把之前某个不重复的数递增到这个数
		如果数X出现了一次以上，就把改数记下来
		如果数X从未出现过，就从记下来的数字中选择一个将其加到X,并对记录做X-i
*/
func minIncrementForUnique1(A []int) int {
	var count [80000]int // 有可能所有的数字39999都要往上加，直接加到80000去了
	for _, v := range A {
		count[v]++
	}


	ans := 0
	op := 0

	for i := range count {
		if count[i] >= 2 {
			// 如果某个数出现了两次及以上
			op += count[i]-1 // 有个多少个重复的数字需要改变
			ans -= i * (count[i] - 1) // 优化： 该数字的值是多少，先欠上，后续找到数字直接还上就可以
		} else if op > 0 && count[i] == 0 {
			// 如果前面有数字需要增加并且当前是个新数字
			op-- // 只改变了前面需要改变的一个数字
			ans += i // 偿还上面优化中欠下的值
		}
	}

	return ans
}

/*
	排序法：先排序，然后遍历，把当前元素每次增加到max+1

*/

func minIncrementForUnique2(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	curMax, ans := -1, 0

	for _, v := range A {
		if v <= curMax {
			// 如果小于当前的最大值， 那么就需要将这个数增加到最大值+1
			ans += curMax+ 1 - v
		}
		curMax = max(v, curMax+1)
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
