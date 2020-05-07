package sliding_window

/*
leet-code 904 水果成篮
链接：https://leetcode-cn.com/problems/fruit-into-baskets

在一排树中，第 i 棵树产生 tree[i] 型的水果。
你可以从你选择的任何树开始，然后重复执行以下步骤：
1. 把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
2. 移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
用这个程序你能收集的水果总量是多少？

示例 1：
输入：[1,2,1]
输出：3
解释：我们可以收集 [1,2,1]。

示例 2：
输入：[0,1,2,2]
输出：3
解释：我们可以收集 [1,2,2].
如果我们从第一棵树开始，我们将只能收集到 [0, 1]。

示例 3：
输入：[1,2,3,2,2]
输出：4
解释：我们可以收集 [2,3,2,2].
如果我们从第一棵树开始，我们将只能收集到 [1, 2]。

示例 4：
输入：[3,3,3,1,2,1,1,2,3,3,4]
输出：5
解释：我们可以收集 [1,2,1,1,2].
如果我们从第一棵树或第八棵树开始，我们将只能收集到 4 个水果。

提示：
1 <= tree.length <= 40000
0 <= tree[i] < tree.length

分析：
 一. 滑动窗口：
    1. 每次只有2种水果
    2. 求2种水果连续最长的长度
*/

// 使用hash
func totalFruit(tree []int) int {
	maxLen := len(tree)
	if maxLen < 2 {
		return maxLen
	}

	var table = make(map[int]int)
	res, l, r := 0, 0, 0
	for ; r < maxLen; r++ {
		table[tree[r]]++
		for len(table) == 3 {
			table[tree[l]]--
			if table[tree[l]] == 0 {
				delete(table, tree[l])
			}
			l++
		}
		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}

// 不使用hash
func totalFruit1(tree []int) int {
	maxLen := len(tree)
	if maxLen < 2 {
		return maxLen
	}

	// k1 for previous and k2 for latter
	k1, k2 := tree[0], -1
	idx := 0
	// initialize k
	for i := 1; i < maxLen; i++ {
		if tree[i] != k1 {
			k2 = tree[i]
			idx = i
			break
		}
	}
	if k2 == -1 || idx == maxLen-1 {
		return maxLen
	}

	res, l, r := idx+1, 0, idx+1
	for ; r < maxLen; r++ {
		if tree[r] == k1 || tree[r] == k2 {
			if r == maxLen-1 && res < r-l+1 {
				res = r - l + 1
				break
			}
			if tree[r] == k1 {
				idx = r
				k1, k2 = k2, k1
			}
			continue
		}

		if res < r-l {
			res = r - l
		}

		l = idx
		idx = r
		k1, k2 = k2, tree[r]
	}

	return res
}
