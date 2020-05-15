package sliding_window

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"0",
		// 	args{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3},
		// 	202,
		// },
		{
			"1",
			args{[]int{11, 49, 100, 20, 86, 29, 72}, 4},
			232,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
