package leetcode

import (
	"reflect"
	"testing"
)

func TestHeightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"height_checker", args{[]int{1, 1, 4, 2, 1, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeightChecker(tt.args.heights); got != tt.want {
				t.Errorf("HeightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnCommonFromSentences(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"uncommon_str1", args{"this apple is sweet", "this apple is sour"}, []string{"sweet", "sour"}},
		{"uncommon_str2", args{"apple apple", "banana"}, []string{"banana"}},
		{"uncommon_str3", args{"d b zu d t", "udb zu ap"}, []string{"b", "t", "udb", "ap"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnCommonFromSentences(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnCommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"data1", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallestFromLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"data0",
			args{
				&TreeNode{0, nil, &TreeNode{1, nil, nil}},
			},
			"ba",
		},
		{
			"data1",
			args{
				&TreeNode{
					25,
					&TreeNode{
						1,
						&TreeNode{1, nil, nil},
						&TreeNode{3, nil, nil},
					},
					&TreeNode{
						3,
						&TreeNode{0, nil, nil},
						&TreeNode{2, nil, nil},
					},
				}},
			"adz",
		},
		{
			"data2",
			args{
				&TreeNode{
					3,
					&TreeNode{9, nil, nil},
					&TreeNode{
						20,
						&TreeNode{15, nil, nil},
						&TreeNode{7, nil, nil},
					},
				},
			},
			"hud",
		},
		{
			"data3",
			args{
				&TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:  1,
						Left: nil,
						Right: &TreeNode{
							Val:  2,
							Left: &TreeNode{6, nil, nil},
							Right: &TreeNode{
								Val:  3,
								Left: nil,
								Right: &TreeNode{
									Val:   4,
									Left:  nil,
									Right: &TreeNode{5, nil, nil},
								},
							},
						},
					},
					Right: nil,
				},
			},
			"fedcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestFromLeaf(tt.args.root); got != tt.want {
				t.Errorf("SmallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"data0",
			args{nums: []int{3, 2, 1, 6, 0, 5}},
			&TreeNode{
				Val:   6,
				Left:  &TreeNode{3, nil, &TreeNode{2, nil, &TreeNode{1, nil, nil}}},
				Right: &TreeNode{5, &TreeNode{0, nil, nil}, nil},
			},
		},
		{"data1",
			args{nums: []int{1, 2, 3}},
			&TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetZeroes(t *testing.T) {
	data := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	printMatrix(data)
	SetZeroes(data)
	printMatrix(data)
}
