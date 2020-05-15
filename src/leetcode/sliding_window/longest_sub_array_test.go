package sliding_window

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{[]int{10, 1, 2, 4, 7, 2}, 5},
			4,
		},
		{
			"1",
			args{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
