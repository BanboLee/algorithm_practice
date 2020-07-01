package graph

/*
leetcode 207. 课程表
链接：https://leetcode-cn.com/problems/course-schedule

你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1:
输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。

提示：
输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
*/

// 拓扑排序： 入度表法
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for i := range adjacency {
		adjacency[i] = make([]int, 0)
	}

	for _, v := range prerequisites {
		inDegrees[v[0]]++
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}

	var queue []int
	for i := range inDegrees {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		numCourses--
		for _, c := range adjacency[cur] {
			inDegrees[c]--
			if inDegrees[c] == 0 {
				queue = append(queue, c)
			}
		}
	}

	return numCourses == 0
}

// 深度优先搜索判断是否有环
func canFinish1(numCourses int, prerequisites [][]int) bool {
	adjacency := make([][]int, numCourses)
	for _, v := range prerequisites {
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}
	flags := make([]int, numCourses)
	for i := range adjacency {
		if !dfs3(&flags, &adjacency, i) {
			return false
		}
	}
	return true
}

func dfs3(flags *[]int, adjacency *[][]int, i int) bool {
	// 本轮已被访问
	if (*flags)[i] == 1 {
		return false
	}
	// 已被其他节点搜索过
	if (*flags)[i] == -1 {
		return true
	}

	(*flags)[i] = 1
	for _, c := range (*adjacency)[i] {
		if !dfs3(flags, adjacency, c) {
			return false
		}
	}
	(*flags)[i] = -1
	return true
}
