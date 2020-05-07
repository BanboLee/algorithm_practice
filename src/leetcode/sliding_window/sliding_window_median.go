package sliding_window

/*
leet-code 480 滑动窗口中位数
链接：https://leetcode-cn.com/problems/sliding-window-median

中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
例如：
[2,3,4]，中位数是 3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，
每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

示例：
给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。
窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
提示：

你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。
与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。

分析：
1. 求中位数，则窗口每次必有序
2. 若每次重新排序，则最低需要O(n*klogk)
3. 使用两个堆，一个大根堆存放一半大的元素，一个小根堆存放一半小的元素
4. 对出窗元素延迟删除，只要不在堆顶，就不会影响结果
5. 每次都要对两个堆进行重新平衡，以免全部倾斜到一边，只能取到最大值
*/

func medianSlidingWindow(nums []int, k int) []float64 {
	minHeap, maxHeap := &heap{}, &heap{order: true}
	// 入窗
	for i := 0; i < k; i++ {
		maxHeap.insert(nums[i]) // 3, 2, 1
	}
	for j := 0; j < k/2; j++ {
		minHeap.insert(maxHeap.pop()) // 4, 5, 6
	}

	hashTable := make(map[int]int)
	var res []float64
	t := k
	for {
		if k%2 != 0 {
			res = append(res, float64(maxHeap.top()))
		} else {
			res = append(res, float64(maxHeap.top()+minHeap.top())*0.5)
		}

		if t >= len(nums) {
			break
		}

		inNum := nums[t]
		outNum := nums[t-k]
		t++
		balance := 0

		// 1. 标记出窗元素
		hashTable[outNum]++
		if outNum <= maxHeap.top() {
			balance++
		} else {
			balance--
		}

		// 2. 入窗
		if inNum > minHeap.top() {
			minHeap.insert(inNum)
			balance++
		} else {
			maxHeap.insert(inNum)
			balance--
		}

		// 3. 重新对大小堆进行平衡
		if balance < 0 {
			minHeap.insert(maxHeap.pop())
			balance++
		} else if balance > 0 {
			maxHeap.insert(minHeap.pop())
			balance--
		}

		// 4.出窗， 当无效元素在堆顶时立即出窗
		for !maxHeap.empty() && hashTable[maxHeap.top()] > 0 {
			hashTable[maxHeap.pop()]--
		}
		for !minHeap.empty() && hashTable[minHeap.top()] > 0 {
			hashTable[minHeap.pop()]--
		}
	}

	return res
}
