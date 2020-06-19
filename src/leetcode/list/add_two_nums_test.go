package list

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := ListFromSlice([]int{2, 4, 3})
	l2 := ListFromSlice([]int{5, 6, 4})
	t.Logf("%v", ListToSlice(addTwoNumbers1(l1, l2)))
	l3 := ListFromSlice([]int{1, 8})
	l4 := ListFromSlice([]int{0})
	t.Logf("%v", ListToSlice(addTwoNumbers1(l3, l4)))

}
