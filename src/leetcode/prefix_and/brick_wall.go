package prefix_and

/*
leetcode 554. 砖墙
链接：https://leetcode-cn.com/problems/brick-wall

你的面前有一堵方形的、由多行砖块组成的砖墙。 这些砖块高度相同但是宽度不同。
你现在要画一条自顶向下的、穿过最少砖块的垂线。砖墙由行的列表表示。
每一行都是一个代表从左至右每块砖的宽度的整数列表。
如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。
你需要找出怎样画才能使这条线穿过的砖块数量最少，并且返回穿过的砖块数量。
你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。

示例：
输入: [[1,2,2,1],
      [3,1,2],
      [1,3,2],
      [2,4],
      [3,1,2],
      [1,3,1,1]]
输出: 2

提示：
每一行砖块的宽度之和应该相等，并且不能超过 INT_MAX。
每一行砖块的数量在 [1,10,000] 范围内， 墙的高度在 [1,10,000] 范围内， 总的砖块数量不超过 20,000。
*/

/*
   题目所求可以理解为每一层相同前缀和的子数组最多的情况
   step1: 计算每一层的前缀和
   step2: 用hash记录每种前缀和出现的次数
   step3: 更新最多次数前缀和
*/
func leastBricks(wall [][]int) int {
	var (
		mp         = make(map[int]int)
		mostPrefix int
	)
	// 直接原地做前缀和计算
	for i := range wall {
		for j := range wall[i] {
			if j != 0 {
				wall[i][j] += wall[i][j-1]
			}

			if j != len(wall[i])-1 {
				mp[wall[i][j]]++
				if mp[wall[i][j]] > mp[mostPrefix] {
					mostPrefix = wall[i][j]
				}
			}

		}
	}
	return len(wall) - mp[mostPrefix]
}
