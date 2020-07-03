package segment_tree

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	st := Constructor([]int{1, 3, 5})
	t.Logf("sum[0,2]=%d", st.SumRange(0, 2))
	st.Update(1, 2)
	t.Logf("sum[0,2]=%d", st.SumRange(0, 2))
}
