package sliding_window

import "testing"

func Test_minKBitFlips(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"data0",
		// 	args{[]int{0, 1, 0}, 1},
		// 	2,
		// },
		// {
		// 	"data1",
		// 	args{[]int{1, 1, 0}, 2},
		// 	-1,
		// },
		{
			"data2",
			args{[]int{0, 1, 0, 0, 1, 0}, 4},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
