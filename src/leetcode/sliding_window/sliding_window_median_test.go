package sliding_window

import (
	"reflect"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"data0",
			args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			[]float64{1, -1, -1, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
