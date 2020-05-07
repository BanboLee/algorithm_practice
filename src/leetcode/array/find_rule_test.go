package array

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"data0",
			args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"data1",
			args{[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"data0",
			args{3},
			[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		args   args
		result [][]int
	}{
		// {
		// 	"data0",
		// 	args{[][]int{
		// 		{1, 2, 3},
		// 		{4, 5, 6},
		// 		{7, 8, 9},
		// 	}},
		// 	[][]int{
		// 		{7, 4, 1},
		// 		{8, 5, 2},
		// 		{9, 6, 3},
		// 	},
		// },
		{
			"data1",
			args{[][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			}},
			[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.matrix)
			for i := range tt.result {
				for j := range tt.result[i] {
					if tt.result[i][j] != tt.args.matrix[i][j] {
						t.Errorf("Rotate() = %v, want %v", tt.args.matrix, tt.result)
					}
				}
			}
		})
	}
}

func TestSearchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"data0",
			args{[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}, 5},
			true,
		},
		{
			"data1",
			args{[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}, 20},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
