package heap

// 实现抽象的优先级队列
type compareFunc func(interface{}, interface{}) bool
type PriorityQueue struct {
	data []interface{}
	cmp  compareFunc
}

func (pq *PriorityQueue) ShiftUp(index int) {
	if index < 0 || index >= len(pq.data) {
		return
	}

	for index > 0 {
		if pq.cmp(pq.data[index], pq.data[(index-1)/2]) {
			pq.data[index], pq.data[(index-1)/2] = pq.data[(index-1)/2], pq.data[index]
			index = (index - 1) / 2
		} else {
			break
		}
	}
}

func (pq *PriorityQueue) Insert(v interface{}) {
	pq.data = append(pq.data, v)
	pq.ShiftUp(len(pq.data) - 1)
}

func (pq *PriorityQueue) ShiftDown(index int) {
	if index < 0 {
		return
	}

	// 需要同时比较两个儿子节点
	for index*2+1 < len(pq.data) {
		t := index*2 + 1
		if t+1 < len(pq.data) && !pq.cmp(pq.data[t], pq.data[t+1]) {
			t++
		}
		if pq.cmp(pq.data[index], pq.data[t]) {
			break
		}
		pq.data[index], pq.data[t] = pq.data[t], pq.data[index]
		index = t
	}
}

func (pq *PriorityQueue) Pop() interface{} {
	if len(pq.data) == 0 {
		return nil
	}
	v := pq.data[0]
	pq.data[0] = pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1]
	pq.ShiftDown(0)
	return v
}

func (pq *PriorityQueue) Top() interface{} {
	if len(pq.data) == 0 {
		return nil
	}
	return pq.data[0]
}

func (pq *PriorityQueue) Size() int {
	return len(pq.data)
}
