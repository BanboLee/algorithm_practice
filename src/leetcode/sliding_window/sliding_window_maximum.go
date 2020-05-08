package sliding_window

/*
leet-code 239 滑动窗口最大值
链接：https://leetcode-cn.com/problems/sliding-window-maximum

给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。



进阶：
你能在线性时间复杂度内解决此题吗？

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

分析：
一、最值问题，维护一个堆就完事了 O(NlogK)
二、双端队列DeQueue O(N)
1. 只需保存每个窗口最大的元素，如果新元素大于这个值，前面的就可以出去了
2. 一直滑动就可以了
3. 从前面出，从后面进
*/

func maxSlidingWindow(nums []int, k int) []int {
	maxHeap := &heap{order: true}
	t := 0
	for ; t < k; t++ {
		maxHeap.insert(nums[t])
	}

	// 维护一个hash表用于延迟删除
	hashTable := make(map[int]int)
	var res []int
	for {
		res = append(res, maxHeap.top())

		if t >= len(nums) {
			break
		}

		inNum := nums[t]
		outNum := nums[t-k]
		t++

		// 标记出窗
		hashTable[outNum]++
		// 入窗
		maxHeap.insert(inNum)
		// 延迟删除
		for hashTable[maxHeap.top()] > 0 {
			hashTable[maxHeap.pop()]--
		}
	}

	return res
}

func maxSlidingWindow1(nums []int, k int) []int {
	dq := &dequeue{}
	// 初始化双向队列
	t := 0
	for ; t < k; t++ {
		for !dq.empty() && nums[t] > nums[dq.head.val] {
			dq.popBack()
		}
		dq.pushBack(t)
	}

	var res []int
	res = append(res, nums[dq.head.val])

	for ; t < len(nums); t++ {
		// 出窗索引小于窗口的
		if !dq.empty() && dq.head.val <= t-k {
			dq.popFront()
		}

		// 如果有个新的最大值，清理小于新值的值
		for !dq.empty() && nums[t] > nums[dq.tail.val] {
			dq.popBack()
		}
		dq.pushBack(t)
		res = append(res, nums[dq.head.val])
	}
	return res
}

// 使用slice作为双端队列
func maxSlidingWindow2(nums []int, k int) []int {
	var (
		res []int
		dq  []int
	)
	for i := 0; i < len(nums); i++ {
		// 出窗
		if len(dq) > 0 && dq[0] == i-k {
			dq = dq[1:]
		}

		// 将小于当前值的索引出队
		for len(dq) > 0 && nums[i] >= nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}

		// 入窗
		dq = append(dq, i)

		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}
	return res
}

// 官方动态规划法。。。。。。。
func maxSlidingWindow3(nums []int, k int) []int {
	var (
		left  = make([]int, len(nums))
		right = make([]int, len(nums))
		res   []int
	)

	for i := 0; i < len(nums); i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}

		j := len(nums) - 1 - i
		if (j+1)%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}

	for t := k - 1; t < len(nums); t++ {
		res = append(res, max(left[t], right[t-k]))
	}

	return res
}
