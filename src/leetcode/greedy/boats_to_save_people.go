package greedy

import "sort"

/*
leetcode 881. 救生艇
链接：https://leetcode-cn.com/problems/boats-to-save-people

第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
返回载到每一个人所需的最小船数。(保证每个人都能被船载)。

示例 1：
输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)
示例 2：
输入：people = [3,2,2,1], limit = 3
输出：3
解释：3 艘船分别载 (1, 2), (2) 和 (3)
示例 3：
输入：people = [3,5,3,4], limit = 5
输出：4
解释：4 艘船分别载 (3), (3), (4), (5)
提示：
1 <= people.length <= 50000
1 <= people[i] <= limit <= 30000
*/

/*
   贪心算法，每次找最重和最轻的人
   最重的人和最轻的人匹配，双指针
   为何不选择最接近limit的，而是最轻的，如果选择了最近接的(比如第二轻的), 那么第一轻的还是会和第二重的匹配，这时依然有空间浪费
   比如limit=7, 直觉上5+2比5+1好, 但是如果5和2能匹配到，2和比5小也更能匹配到，一样找得到船
*/
func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] < people[j]
	})
	light, heavy := 0, len(people)-1
	ans := 0
	for light <= heavy {
		ans++
		if people[heavy]+people[light] <= limit {
			light++
		}
		heavy--
	}
	return ans
}
