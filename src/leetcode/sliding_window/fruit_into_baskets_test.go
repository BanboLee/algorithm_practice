package sliding_window

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"data-2",
			args{[]int{0, 1, 6, 6, 4, 4, 6}},
			5,
		},
		{
			"data-1",
			args{[]int{1, 2, 1}},
			3,
		},
		{
			"data0",
			args{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit1(tt.args.tree); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
