package greedy

import "sort"

/*
leetcode 406. 根据身高重建队列
链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height

假设有打乱顺序的一群人站成一个队列。
每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。

注意：
总人数少于1100人。

示例
输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/

/*
   按照身高降序和K值升序排序，然后将各个数据插入到对应的K值位置
*/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	var output [][]int
	for _, p := range people {
		k := p[1]
		var tmp [][]int
		tmp = append(tmp, output[:k]...)
		tmp = append(tmp, p)
		output = append(tmp, output[k:]...)
	}

	return output
}
