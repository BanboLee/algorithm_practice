package algorithm

func ShortestPathDijkstra(g [][]int, t int, max int) []int {
	var result = make([]int, len(g))
	for i := range result {
		if i == t {
			result[i] = 0
			continue
		}
		result[i] = max
	}
	var q = make([]int, len(g))
	for i := range g {
		q[i] = 0
	}

	for k := 0; k < len(q); k++ {
		// extract min from q.
		shortest, v := max+1, -1
		for i, isVis := range q {
			if isVis == 1 {
				continue
			}
			if result[i] < shortest {
				shortest = result[i]
				v = i
			}
		}
		q[v] = 1

		// relaxation.
		for i, d := range g[v] {
			if result[i] > d+shortest {
				result[i] = d + shortest
			}
		}
	}

	return result
}
