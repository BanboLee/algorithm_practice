package list

import (
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	list0 := &ListNode{1,
		&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list1 := oddEvenList(list0)
	t.Logf("%v\n", ListToSlice(list1))
}
