package heap

import "testing"

func TestHeap(t *testing.T) {
	testData := []int{8, 5, 2, 10, 3, 7, 1, 4, 6}
	hp := &heap{}
	hp.fromSlice(testData)
	t.Log(hp.data, hp.pop(), hp.pop())
}
