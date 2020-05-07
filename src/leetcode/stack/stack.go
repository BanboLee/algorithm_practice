package stack

import "fmt"

type stack struct {
	data []int
}

func (s *stack) push(a int) {
	s.data = append(s.data, a)
}

func (s *stack) pop() int {
	if len(s.data) == 0 {
		return -1
	}
	a := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return a
}

// 括号生成
func GenerateBrackets(n int) []string {
	var result []string
	bracketsMatch(1, 0, "(", n, &result)
	return result
}

func bracketsMatch(lh int, rh int, brackets string, n int, result *[]string) {
	if lh < n {
		bracketsMatch(lh+1, rh, brackets+"(", n, result)
	}
	if rh < lh {
		bracketsMatch(lh, rh+1, brackets+")", n, result)
	}

	if rh == n && lh == n {
		*result = append(*result, brackets)
	}
}

// 布尔运算
var boolMap = make(map[string][2]int)

func CountBoolEval(s string, result int) int {
	if s == "" {
		return 0
	}
	computeBoolExpr(s)

	return boolMap[s][result]
}

func computeBoolExpr(s string) {
	if len(s) == 1 {
		if s == "1" {
			boolMap[s] = [2]int{0, 1}
		} else {
			boolMap[s] = [2]int{1, 0}
		}
		return
	}
	var ans = [2]int{0, 0}

	for idx, c := range s {
		if idx%2 == 0 {
			continue
		}
		left, leftOk := boolMap[s[:idx]]
		right, rightOk := boolMap[s[idx+1:]]
		var left0 int
		var left1 int
		var right0 int
		var right1 int
		if leftOk {
			left0, left1 = left[0], left[1]
		} else {
			computeBoolExpr(s[:idx])
			left0, left1 = boolMap[s[:idx]][0], boolMap[s[:idx]][1]
		}
		if rightOk {
			right0, right1 = right[0], right[1]
		} else {
			computeBoolExpr(s[idx+1:])
			right0, right1 = boolMap[s[idx+1:]][0], boolMap[s[idx+1:]][1]
		}

		switch c {
		case '&':
			ans[0] += left1*right0 + left0*right1 + left0*right0
			ans[1] += left1 * right1
		case '|':
			ans[0] += left0 * right0
			ans[1] += left1*right0 + left0*right1 + left1*right1
		case '^':
			ans[0] += left1*right1 + left0*right0
			ans[1] += left1*right0 + left0*right1
		}

		fmt.Printf("s=%s, ans=%v\n", s, ans)
		boolMap[s] = ans
	}
}

// 用两个栈实现队列
type CQueue struct {
	stack1 *stack
	stack2 *stack
}

func Constructor() CQueue {
	return CQueue{&stack{}, &stack{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.push(value)
}

func (this *CQueue) DeleteHead() int {
	a := this.stack1.pop()
	if a == -1 {
		return -1
	}
	for a != -1 {
		this.stack2.push(a)
		a = this.stack1.pop()
	}
	b := this.stack2.pop()
	a = this.stack2.pop()
	for a != -1 {
		this.stack1.push(a)
		a = this.stack2.pop()
	}
	return b
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	left, right := 0, k
	maxIdx := -1
	var result []int
	for i := 0; i < len(nums)-k+1; i++ {
		if maxIdx < left {
			maxIdx = left
			for j := left; j < right; j++ {
				if nums[j] > nums[maxIdx] {
					maxIdx = j
				}
			}
		} else if nums[right-1] > nums[maxIdx] {
			maxIdx = right - 1
		}

		result = append(result, nums[maxIdx])
		left++
		right++
	}
	return result
}
