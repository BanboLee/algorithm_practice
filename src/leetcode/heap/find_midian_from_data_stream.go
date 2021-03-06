package heap

import "fmt"

/*
leetcode 295. 数据流的中位数
链接：https://leetcode-cn.com/problems/find-median-from-data-stream

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例：
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:
如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
*/

/*
   使用两个堆来操作, 一个存放一半比较大的数据，另一个存放比较小的数据
   最大堆比最小堆可以多存放一个元素
*/
type MedianFinder struct {
	lowerHeap  *PriorityQueue // 存放小的一半数据
	higherHeap *PriorityQueue // 存放大的一半数据
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		higherHeap: &PriorityQueue{nil, func(i interface{}, i2 interface{}) bool {
			return i.(int) < i2.(int)
		}},
		lowerHeap: &PriorityQueue{nil, func(i interface{}, i2 interface{}) bool {
			return i.(int) > i2.(int)
		}},
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.lowerHeap.Insert(num)

	this.higherHeap.Insert(this.lowerHeap.Pop())
	if this.lowerHeap.Size() < this.higherHeap.Size() {
		this.lowerHeap.Insert(this.higherHeap.Pop())
	}

	fmt.Printf("after add %d, lowerHeap=%v, higher_heap=%v\n", num, this.lowerHeap.data, this.higherHeap.data)
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lowerHeap.Size() == 0 && this.higherHeap.Size() == 0 {
		return 0
	}
	if this.lowerHeap.Size() == this.higherHeap.Size() {
		return (float64(this.higherHeap.Top().(int)) + float64(this.lowerHeap.Top().(int))) * 0.5
	}
	return float64(this.lowerHeap.Top().(int))
}
