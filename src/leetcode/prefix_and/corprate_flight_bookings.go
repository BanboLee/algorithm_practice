package prefix_and

/*
leetcode 1109. 航班预订统计
链接：https://leetcode-cn.com/problems/corporate-flight-bookings

这里有 n 个航班，它们分别从 1 到 n 进行编号。
我们这儿有一份航班预订表，表中第 i 条预订记录 bookings[i] = [i, j, k] 意味着我们在从 i 到 j 的每个航班上预订了 k 个座位。
请你返回一个长度为 n 的数组 answer，按航班编号顺序返回每个航班上预订的座位数。

示例：
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]

提示：
1 <= bookings.length <= 20000
1 <= bookings[i][0] <= bookings[i][1] <= n <= 20000
1 <= bookings[i][2] <= 10000
*/

// 普通方法 最坏情况O(n²)
func corpFlightBookings(bookings [][]int, n int) []int {
	var (
		seats [20000]int
	)

	for _, b := range bookings {
		for i := b[0]; i <= b[1]; i++ {
			seats[i-1] += b[2]
		}
	}

	return seats[:n]
}

/*
   上下车问题(差分思想)
*/
func corpFlightBookings1(bookings [][]int, n int) []int {
	// 记录每一站人数变化
	var changes = make([]int, n)
	for _, b := range bookings {
		// 上车
		if b[0]-1 < n {
			changes[b[0]-1] += b[2]
		}

		// 下车
		if b[1] < n {
			changes[b[1]] -= b[2]
		}
	}

	// 前缀和即为每站剩余人数
	for i := 1; i < n; i++ {
		changes[i] += changes[i-1]
	}

	return changes
}
