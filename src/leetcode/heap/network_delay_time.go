package heap

import (
	"leetcode/algorithm"
)

func networkDelayTime(times [][]int, N int, K int) int {
	// construct graph
	var data = make([][]int, N)
	for i := range data {
		data[i] = make([]int, N)
		for j := range data[i] {
			if i == j {
				data[i][j] = 0
			} else {
				data[i][j] = N * 6000
			}
		}
	}
	for _, t := range times {
		cell0 := t[0] - 1
		cell1 := t[1] - 1
		data[cell0][cell1] = t[2]
	}

	allTime := algorithm.ShortestPathDijkstra(data, K-1, N*6000)
	// allTime := algorithm.ShortestPathBellmanFord(data, K-1, N*6000)
	maxTime := -1
	for _, v := range allTime {
		if v == N*6000 {
			maxTime = -1
			break
		}
		if v > maxTime {
			maxTime = v
		}
	}
	return maxTime
}
