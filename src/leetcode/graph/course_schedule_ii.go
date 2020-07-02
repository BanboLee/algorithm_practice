package graph

/*
leetcode 210. 课程表 II
链接：https://leetcode-cn.com/problems/course-schedule-ii

现在你总共有 n 门课需要选，记为 0 到 n-1。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

示例 1:
输入: 2, [[1,0]]
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2:
输入: 4, [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
说明:
输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
*/

// 广度优先搜索(入度表)
func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int
	inDegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for _, v := range prerequisites {
		inDegrees[v[0]]++
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}

	var deQueue []int
	for i := range inDegrees {
		if inDegrees[i] == 0 {
			deQueue = append(deQueue, i)
		}
	}

	for len(deQueue) != 0 {
		cur := deQueue[0]
		deQueue = deQueue[1:]
		numCourses--
		res = append(res, cur)

		for _, c := range adjacency[cur] {
			inDegrees[c]--
			if inDegrees[c] == 0 {
				deQueue = append(deQueue, c)
			}
		}
	}

	if numCourses == 0 {
		return res
	}

	return nil
}

// 深度优先搜索实现拓扑排序
func dfs4(flags *[]int, adjacency *[][]int, i int, res *[]int) bool {
	// 本轮已被访问，有环
	if (*flags)[i] == 1 {
		return false
	}
	// 已被其他节点启动
	if (*flags)[i] == -1 {
		return true
	}
	(*flags)[i] = 1
	for _, c := range (*adjacency)[i] {
		if !dfs4(flags, adjacency, c, res) {
			return false
		}
	}
	(*flags)[i] = -1
	*res = append(*res, i)
	return true
}

func findOrder1(numCourses int, prerequisites [][]int) []int {
	var res []int
	flags := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for _, v := range prerequisites {
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}

	for i := range adjacency {
		if !dfs4(&flags, &adjacency, i, &res) {
			return nil
		}
	}

	return reverseIntSlice(res)
}

func reverseIntSlice(data []int) []int {
	for i := 0; i <= len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	return data
}