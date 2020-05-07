package sliding_window

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"data0",
			args{[]int{8, 8, 9, 7, 10, 9, 9, 1, 9}},
			5,
		},
		{
			"data1",
			args{[]int{4, 8, 12, 16}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.A); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
