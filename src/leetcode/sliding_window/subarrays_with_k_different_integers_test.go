package sliding_window

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"data0",
			args{[]int{1, 2, 1, 3, 4}, 3},
			3,
		},
		{
			"data1",
			args{[]int{1, 2, 1, 2, 3}, 2},
			7,
		},
		{
			"data2",
			args{[]int{1, 2}, 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct1(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
