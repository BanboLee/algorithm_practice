package list

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := ListFromSlice([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := ListFromSlice([]int{5, 6, 4})
	t.Logf("%v", ListToSlice(addTwoNumbers(l1, l2)))
}
