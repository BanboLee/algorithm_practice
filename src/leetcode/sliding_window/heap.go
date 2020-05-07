package sliding_window

import (
	"math"
)

type heap struct {
	data  []int
	order bool // true for positive and false for negative
}

func (h *heap) shiftUp(i int) {
	if i <= 0 || i >= len(h.data) {
		return
	}

	for i > 0 {
		p := (i - 1) / 2
		if (h.order && h.data[p] < h.data[i]) ||
			(!h.order && h.data[p] > h.data[i]) {
			h.data[p], h.data[i] = h.data[i], h.data[p]
			i = p
		} else {
			break
		}
	}
}

func (h *heap) shiftDown(i int) {
	l, r := 2*i+1, 2*i+2
	if i < 0 || (l >= len(h.data) && r >= len(h.data)) {
		return
	}

	if l < len(h.data) && r < len(h.data) {
		if (h.data[l] < h.data[r] && h.order) ||
			(h.data[l] > h.data[r] && !h.order) {
			l, r = r, l
		}
	}
	if l >= len(h.data) {
		l = r
	}
	if (h.order && h.data[i] < h.data[l]) ||
		(!h.order && h.data[i] > h.data[l]) {
		h.data[i], h.data[l] = h.data[l], h.data[i]
		h.shiftDown(l)
	}
}

func (h *heap) insert(v int) {
	h.data = append(h.data, v)
	h.shiftUp(len(h.data) - 1)
}

func (h *heap) pop() int {
	if len(h.data) == 0 {
		return math.MinInt32
	}
	res := h.data[0]
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[:len(h.data)-1]
	h.shiftDown(0)
	return res
}

func (h *heap) top() int {
	if len(h.data) == 0 {
		return math.MinInt32
	}
	return h.data[0]
}

func (h *heap) empty() bool {
	return len(h.data) == 0
}
