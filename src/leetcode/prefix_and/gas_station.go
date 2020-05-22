package prefix_and

import "math"

/*
leetcode 134. 加油站
链接：https://leetcode-cn.com/problems/gas-station

在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

说明:
如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。

示例 1:
输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

示例 2:
输入:
gas  = [2,3,4]
cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。
*/

/*
   官方题解：
        条件1: sum(gas) > sum(cost), 总的汽油量一定大于总的消耗量
        条件2: gas[i] < cost[i], 则不可能从i出发
    思考：
        step1: 从头遍历某个加油站，检查从这个加油站跑一圈能否到达
        step2: 优化step1, 如果某个加油站不满足条件2, 直接跳过
        step3: 如果从start无法到达当前加油站i, 即不满足条件1
             3.1. 起点不会是start左边, 因为从step2切换为了当前start, 不然不会发生切换起点
             3.2. 起点不会是(start, i)中的start`, 在此条件下
                  如果start无法到达start`就会发生切换，因此start`无法到达i
                  但是这种情况下，(start`, i]的总剩余油量sum必然为负， 因此起点必不在i左边
        step4: 起点在i右边，直至全部遍历完成，sum > 0则找到起点， sum < 0 则无解
    复杂度：时间O(N), 空间O(1)
*/
func canCompleteCircuit(gas []int, cost []int) int {
	var (
		res        int
		remainTank int
		curTank    int
	)

	for i := 0; i < len(gas); i++ {
		remainTank += gas[i] - cost[i]
		curTank += gas[i] - cost[i]

		if curTank < 0 {
			res = i + 1 // 从0开始遍历的，不用取模回到0
			curTank = 0
		}
	}

	if remainTank < 0 {
		return -1
	}

	return res
}

/*
   来自大佬的折线图解法：
   step1: 将到达每个站点的剩余油量画成折线图
   step2: 折线再X轴下方的站点，无法作为起点，在上方才可以
   step3: 向上移动折线图，直到整个折线图都在X轴上方
   step4: 题目具有唯一解，因此会有一个点正好在X轴上为起点
*/
func canCompleteCircuit1(gas []int, cost []int) int {
	var (
		res           int
		remainTank    int             // 从起点至当前的剩余油量
		minRemainTank = math.MaxInt32 // 最低剩余油量
	)

	for i := 0; i < len(gas); i++ {
		remainTank += gas[i] - cost[i]
		// 直到到达那个剩余油量开始上升的地点以后, remainTank >=  minRemainTank 一直成立
		// 这个地点就是折线全部上升到X轴上方后的那个最低点
		if remainTank < minRemainTank {
			minRemainTank = remainTank
			res = i
		}
	}

	if remainTank < 0 {
		return -1
	}

	// 取模获得起点
	return (res + 1) % len(gas)
}
