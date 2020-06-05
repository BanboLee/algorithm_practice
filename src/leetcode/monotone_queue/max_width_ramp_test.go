package monotone_queue

import "testing"

func Test_maxWidthRamp(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{[]int{6, 0, 8, 2, 1, 5}},
			4,
		},
		{
			"1",
			args{[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.args.A); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
