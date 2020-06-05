package list

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	data0 := []int{1, 2, 3, 4, 5}
	t.Logf("data0=%v", data0)
	head0 := ListFromSlice(data0)
	t.Logf("list0=%v", ListToSlice(head0))
	head1 := reverseList(head0)
	t.Logf("list1=%v", ListToSlice(head1))
}
