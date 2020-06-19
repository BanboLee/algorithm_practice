package tree

import (
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	l := []int{1, 2}
	mid := len(l) / 2

	t.Logf("%v, %v", l[:mid], l[mid+1:])

}
