package sliding_window

/*
leetcode 面试题59 - II. 队列的最大值
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof

请定义一个队列并实现函数 max_value 得到队列里的最大值，
要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：
输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]

示例 2：
输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

限制：
1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/

type MaxQueue struct {
	data   []int
	dQueue []int // 双端队列
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.data) > 0 {
		return this.dQueue[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.dQueue) > 0 && value > this.dQueue[len(this.dQueue)-1] {
		this.dQueue = this.dQueue[:len(this.dQueue)-1]
	}

	this.dQueue = append(this.dQueue, value)
	this.data = append(this.data, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	v := this.data[0]
	this.data = this.data[1:]
	if this.dQueue[0] == v {
		this.dQueue = this.dQueue[1:]
	}
	return v
}
