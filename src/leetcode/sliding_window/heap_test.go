package sliding_window

import "testing"

func TestHeap(t *testing.T) {
}

func Test_heap_empty(t *testing.T) {
	type fields struct {
		data  []int
		order bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data:  tt.fields.data,
				order: tt.fields.order,
			}
			if got := h.empty(); got != tt.want {
				t.Errorf("empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_insert(t *testing.T) {
	type fields struct {
		data  []int
		order bool
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"data0",
			fields{[]int{1, 3, -1}, true},
			args{-3},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				order: tt.fields.order,
			}
			for _, d := range tt.fields.data {
				h.insert(d)
				t.Logf("%v", h.data)
			}
			h.insert(tt.args.v)
			if h.data[0] != tt.want {
				t.Errorf("want top()=%v, got %v", tt.want, h.data[0])
			}
		})
	}
}

func Test_heap_pop(t *testing.T) {
	type fields struct {
		data  []int
		order bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				order: tt.fields.order,
			}

			if got := h.pop(); got != tt.want {
				t.Errorf("pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_shiftDown(t *testing.T) {
	type fields struct {
		data  []int
		order bool
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// h := &heap{
			// 	data:  tt.fields.data,
			// 	order: tt.fields.order,
			// }
		})
	}
}

func Test_heap_shiftUp(t *testing.T) {
	type fields struct {
		data  []int
		order bool
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// h := &heap{
			// 	data:  tt.fields.data,
			// 	order: tt.fields.order,
			// }
		})
	}
}

func Test_heap_top(t *testing.T) {
	type fields struct {
		data  []int
		order bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data:  tt.fields.data,
				order: tt.fields.order,
			}
			if got := h.top(); got != tt.want {
				t.Errorf("top() = %v, want %v", got, tt.want)
			}
		})
	}
}
