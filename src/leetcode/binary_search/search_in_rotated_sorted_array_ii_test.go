package binary_search

import "testing"

func Test_search1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"0",
			args{[]int{1, 1, 3, 1}, 3},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search1() = %v, want %v", got, tt.want)
			}
		})
	}
}
