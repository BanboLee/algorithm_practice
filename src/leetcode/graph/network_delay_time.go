package graph

import "math"

/*
leetcode 743. 网络延迟时间
链接：https://leetcode-cn.com/problems/network-delay-time

有 N 个网络节点，标记为 1 到 N。
给定一个列表 times，表示信号经过有向边的传递时间。
times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。
现在，我们从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。

示例：
输入：times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
输出：2

注意:
N 的范围在 [1, 100] 之间。
K 的范围在 [1, N] 之间。
times 的长度在 [1, 6000] 之间。
所有的边 times[i] = (u, v, w) 都有 1 <= u, v <= N 且 0 <= w <= 100。
*/

// dijkstra
func networkDelayTime(times [][]int, N int, K int) int {
	// 邻接矩阵存储图
	graph := make([][]int, N)
	for i := range graph {
		for j := 0; j < N; j++ {
			if i == j {
				graph[i] = append(graph[i], 0)
			} else {
				graph[i] = append(graph[i], math.MaxInt32)
			}
		}
	}
	for _, v := range times {
		graph[v[0]-1][v[1]-1] = v[2]
	}

	dist := make([]int, N)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[K-1] = 0
	st := make([]bool, N) // 是否已找到最优解

	for i := 0; i < N; i++ {
		// 从还未找到最短距离的点中找到当前最短的点
		t := -1
		for j := 0; j < N; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		st[t] = true // t点的最短路径已确定
		// 使用t更新其他点的距离
		for j := 0; j < N; j++ {
			dist[j] = min(dist[j], dist[t]+graph[t][j])
		}
	}

	maxTime := 0
	for _, v := range dist {
		if v > maxTime {
			maxTime = v
		}
	}

	if maxTime == math.MaxInt32 {
		return -1
	}
	return maxTime
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
