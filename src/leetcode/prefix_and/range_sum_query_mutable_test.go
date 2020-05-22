package prefix_and

import "testing"

func TestNumArray1_SumRange1(t *testing.T) {
	na := Constructor2([]int{1, 3, 5})
	t.Log(na.SumRange1(0, 2))
	na.Update(1, 2)

}
