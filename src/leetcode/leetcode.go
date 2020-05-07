package leetcode

import (
	"fmt"
	"strings"
)

func HeightChecker(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	var dep = make(map[int]int)

	cnt := 0
	for i := range heights {
		idx := 0
		for j := range heights[:len(heights)-i] {
			if heights[idx] < heights[j] {
				idx = j
			}
		}

		if idx != len(heights)-i-1 && heights[idx] != heights[len(heights)-i-1] {
			heights[idx], heights[len(heights)-i-1] = heights[len(heights)-i-1], heights[idx]
			if dep[idx] != 1 {
				dep[idx] = 1
				cnt++
			}
			if dep[len(heights)-i-1] != 1 {
				dep[len(heights)-i-1] = 1
				cnt++
			}
		}
	}

	return cnt
}

func UnCommonFromSentences(a, b string) []string {
	var ret []string
	as := strings.Split(a, " ")
	for _, s := range as {
		prefix, mid, suffix := s+" ", " "+s+" ", " "+s
		if !(strings.HasPrefix(a, prefix) && strings.Contains(a, mid)) &&
			!(strings.HasSuffix(a, suffix) && strings.Contains(a, mid)) &&
			!(strings.HasPrefix(a, prefix) && strings.HasSuffix(a, suffix)) &&
			(strings.Index(a, mid) < 0 || strings.Index(a, mid) == strings.LastIndex(a, mid)) &&
			!strings.HasPrefix(b, prefix) && !strings.HasSuffix(b, suffix) && !strings.Contains(b, mid) {
			ret = append(ret, s)
		}
	}

	bs := strings.Split(b, " ")
	for _, s := range bs {
		prefix, mid, suffix := s+" ", " "+s+" ", " "+s
		if !(strings.HasPrefix(b, prefix) && strings.Contains(b, mid)) &&
			!(strings.HasSuffix(b, suffix) && strings.Contains(b, mid)) &&
			!(strings.HasPrefix(b, prefix) && strings.HasSuffix(b, suffix)) &&
			(strings.Index(b, mid) < 0 || strings.Index(b, mid) == strings.LastIndex(b, mid)) &&
			!strings.HasPrefix(a, prefix) && !strings.HasSuffix(a, suffix) && !strings.Contains(a, mid) {
			ret = append(ret, s)
		}
	}

	return ret
}

func FindDisappearedNumbers(nums []int) []int {
	var ret []int

	for i := range nums {
		next := nums[i]
		if next == -1 {
			continue
		}
		for nums[next-1] != -1 {
			last := next
			next = nums[next-1]
			nums[last-1] = -1
		}
	}

	for j, v := range nums {
		if v != -1 {
			ret = append(ret, j+1)
		}
	}

	return ret
}

/*
func Intersect(nums1 []int, nums2 []int) []int {
	var ret []int

	for _, v1 := range nums1 {
		for i, v2 := range nums2 {
			if v1 == v2 {
				ret = append(ret, v1)
				nums2[i] =
			}
		}
	}
	return ret
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPath(root *TreeNode, tail string) string {
	if root.Left == nil && root.Right == nil {
		return string(root.Val+97) + tail
	}
	if root.Left == nil {
		return findPath(root.Right, string(root.Val+97)+tail)
	}
	if root.Right == nil {
		return findPath(root.Left, string(root.Val+97)+tail)
	}

	ls, rs := findPath(root.Left, string(root.Val+97)+tail), findPath(root.Right, string(root.Val+97)+tail)
	if ls > rs {
		return rs
	}
	return ls
}

func SmallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}

	ret := findPath(root, "")
	return ret
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	idx, max := 0, nums[0]
	for i, v := range nums {
		if v > max {
			idx = i
			max = v
		}
	}
	fmt.Printf("nums=%v, idx=%d, max=%d\n", nums, idx, max)
	return &TreeNode{
		Val:   max,
		Left:  ConstructMaximumBinaryTree(nums[:idx]),
		Right: ConstructMaximumBinaryTree(nums[idx+1:]),
	}
}

func printMatrix(matrix [][]int) {
	fmt.Println("----------------------------")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
	fmt.Println("----------------------------")
}

func SetZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	row0NeedSet, col0NeedSet := false, false
	// 计算第一行
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			row0NeedSet = true
			break
		}
	}
	// 计算第一列
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0NeedSet = true
			break
		}
	}

	// 先计算除第一行和第一列的数据
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				// 将该位置左边和上边的元素都设置为0
				for k := j - 1; k >= 0; k-- {
					matrix[i][k] = 0
				}
				for k := i - 1; k >= 0; k-- {
					matrix[k][j] = 0
				}
			} else {
				// 如果他左边和上边都为0,那该元素也设置为0
				allZero := true
				for k := j - 1; k >= 0; k-- {
					if matrix[i][k] != 0 {
						allZero = false
						break
					}
				}
				if j > 0 && allZero {
					matrix[i][j] = 0
					continue
				}
				allZero = true
				for k := i - 1; k >= 0; k-- {
					if matrix[k][j] != 0 {
						allZero = false
						break
					}
				}
				if i > 0 && allZero {
					matrix[i][j] = 0
				}
			}
		}
		fmt.Printf("row=%d, data=%v\n", i, matrix[i])
	}

	if row0NeedSet {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
	if col0NeedSet {
		for j := range matrix {
			matrix[j][0] = 0
		}
	}
}
