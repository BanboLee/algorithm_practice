package sliding_window

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow1(t *testing.T) {
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
			"data0",
			args{[]int{1, 3, 1, 2, 0, 5}, 3},
			[]int{3, 3, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow1(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow1() = %v, want %v", got, tt.want)
			}
		})
	}
}
