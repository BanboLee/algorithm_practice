package stack

import "testing"

func TestGenerateBrackets(t *testing.T) {
	t.Logf("%v", GenerateBrackets(3))
}

func TestCountBoolEval(t *testing.T) {
	type args struct {
		s      string
		result int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{s: "1^0|0|1", result: 0},
			2,
		},
		{
			"1",
			args{s: "0&0&0&1^1|0", result: 1},
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBoolEval(tt.args.s, tt.args.result); got != tt.want {
				t.Errorf("CountBoolEval() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestOthers(t *testing.T) {
	const sample = "A"
	t.Log(sample)
	for _, b := range sample {
		t.Logf("%x", b)
	}
	t.Logf("% x", sample)
	t.Logf("1&1=%d,1&0=%d, 0|1=%d", 1&1, 1&0, 0|1)

	s := "0&0&0&1^1|0"
	for idx, c := range s {
		if idx%2 == 0 {
			continue
		}
		t.Logf("idx=%d, c=%c", idx, c)
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			"2",
			args{[]int{7, 2, 4}, 2},
			[]int{7, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSlidingWindow(tt.args.nums, tt.args.k)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("maxSlidingWindow = %v, want %v", got, tt.want)
				}
			}
			if len(got) != len(tt.want) {
				t.Errorf("maxSlidingWindow = %v, want %v", got, tt.want)
			}
		})
	}
}
