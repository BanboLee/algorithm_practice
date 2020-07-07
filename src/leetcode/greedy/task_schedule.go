package greedy

import "sort"

/*
leetcode 621. 任务调度器
链接：https://leetcode-cn.com/problems/task-scheduler

给定一个用字符数组表示的 CPU 需要执行的任务列表。
其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。
任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。
CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
你需要计算完成所有任务所需要的最短时间。

示例 ：
输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
提示：
任务的总个数为 [1, 10000]。
n 的取值范围为 [0, 100]。
*/

/*
   选择最多的任务p作为第一个任务，将需要等待的(p-1)*n个空闲位，然后填入其他任务
   如果空闲位有剩余，则空闲的位置会空转，结果即为空闲位数量+任务数量
   如果没有空闲位，则没有空转的可能，即所有任务数量为结果
*/
func leastInterval(tasks []byte, n int) int {
	var mp [26]int
	for _, ch := range tasks {
		mp[ch-'A']++
	}

	sort.Slice(mp[:], func(i, j int) bool {
		if mp[i] > mp[j] {
			return true
		}
		return false
	})

	nMax := mp[25] - 1
	idle := nMax * n // 空闲位置
	for i := 24; i >= 0; i-- {
		// 这里如果某个任务数量和最大任务数相同时，仍让其占用nMax个位置
		// 相当于多的这一个任务填到最后一行去了， 因为如果所有任务占满了空闲，结果都是任务总数
		// 而如果没有占满，那么剩余的空闲数依然可以和这个位置折中合并
		idle -= min(nMax, mp[i])
	}

	if idle > 0 {
		return idle + len(tasks)
	}

	return len(tasks)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
