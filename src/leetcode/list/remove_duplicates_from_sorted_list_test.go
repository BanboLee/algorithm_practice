package list

import (
	"testing"
)

func Test_deleteDuplicates1(t *testing.T) {
	list0 := &ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	t.Logf("%v", ListToSlice(deleteDuplicates1(list0)))
}
