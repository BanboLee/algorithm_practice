package array

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"data0",
			args{[]int{0, 1, 2, 3, 4, 5}, 4},
		},
		{
			"data1",
			args{[]int{1, 2}, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v", tt.args.nums)
			rotate(tt.args.nums, tt.args.k)
			t.Logf("%v", tt.args.nums)
		})
	}
}

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"data0",
			args{[]int{3, 2, 1}},
			[]int{1, 2, 3},
		},
		{
			"data1",
			args{[]int{1, 3, 2}},
			[]int{2, 1, 3},
		},
		{
			"data2",
			args{[]int{1, 1, 5}},
			[]int{1, 5, 1},
		},
		{
			"data3",
			args{[]int{1, 2, 3}},
			[]int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("input=%v", tt.args.nums)
			nextPermutation(tt.args.nums)
			for i, v := range tt.want {
				if tt.args.nums[i] != v {
					t.Errorf("\nwant=%v\ngot=%v", tt.want, tt.args.nums)
					break
				}
			}
		})
	}
}

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"data0",
			args{[]int{1, 3, 4, 2, 2}},
			2,
		},
		{
			"data1",
			args{[]int{3, 1, 3, 4, 2}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate2(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"data0",
			args{[]int{1, 2, 0}},
			3,
		},
		{
			"data1",
			args{[]int{3, 4, -1, 1}},
			2,
		},
		{
			"data2",
			args{[]int{7, 8, 9, 11, 12}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
