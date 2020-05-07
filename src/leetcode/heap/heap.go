package heap

import "fmt"

type heap struct {
	data []int
}

func (h *heap) fromSlice(data []int) {
	h.data = data
	h.minHeapify()
}

func (h *heap) shiftUp(index int) {
	if index <= 0 || index >= len(h.data) {
		return
	}

	for index > 0 {
		if h.data[index] < h.data[(index-1)/2] {
			h.data[index], h.data[(index-1)/2] = h.data[(index-1)/2], h.data[index]
			index = (index - 1) / 2
		}
	}
}

func (h *heap) shiftDown(index int) {
	if index < 0 || index > (len(h.data)-1)/2 {
		return
	}

	// fmt.Println("==================================")
	// fmt.Printf("shift down %d: \n", index)
	// h.print()
	for index*2+1 < len(h.data) {
		li, ri := index*2+1, index*2+2
		if li < len(h.data) && ri < len(h.data) && h.data[li] > h.data[ri] {
			li, ri = ri, li
		}
		if h.data[index] > h.data[li] {
			h.data[index], h.data[li] = h.data[li], h.data[index]
			index = li
		} else {
			break
		}
	}
	//
	// fmt.Printf("after shift down: \n")
	// h.print()
}

func (h *heap) insert(v int) {
	h.data = append(h.data, v)
	h.shiftUp(len(h.data) - 1)
}

func (h *heap) pop() int {
	if len(h.data) == 0 {
		return -1
	}

	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	value := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.shiftDown(0)
	return value
}

func (h *heap) top() int {
	if len(h.data) == 0 {
		return -1
	}

	return h.data[0]
}

func (h *heap) minHeapify() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *heap) print() {
	start := 1
	cnt := 0
	for k := 0; k < (len(h.data)+1)/2; k++ {
		fmt.Print(" ")
	}
	for i := 0; i < len(h.data); i++ {
		if cnt == start {
			start += cnt
			cnt = 0
			fmt.Println()
			for j := 0; j < (len(h.data)-start)/2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Printf(" %d-%d", i, h.data[i])
		cnt++
	}
	fmt.Println()
}
