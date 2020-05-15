package sliding_window

import (
	"reflect"
	"testing"
)

func Test_numMovesStonesII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"data0",
			args{[]int{8, 7, 6, 5, 10}},
			[]int{1, 1},
		},
		{
			"data1",
			args{[]int{7, 4, 9}},
			[]int{1, 2},
		},
		{
			"data2",
			args{[]int{100, 101, 104, 102, 103}},
			[]int{0, 0},
		},
		{
			"data3",
			args{[]int{1, 500000000, 1000000000}},
			[]int{2, 499999999},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStonesII(tt.args.stones); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStonesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
