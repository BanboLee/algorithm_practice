package heap

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"data0",
			args{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2},
			2,
		},
		{"data1",
			args{[][]int{{1, 2, 1}}, 2, 2},
			-1,
		},
		{
			"data2",
			args{[][]int{{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30}, {2, 1, 47}, {1, 5, 61}, {1, 4, 99},
				{3, 4, 68}, {3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44}, {4, 5, 19}, {2, 3, 13}, {3, 2, 18},
				{1, 2, 0}, {5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74}}, 5, 3},
			59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
