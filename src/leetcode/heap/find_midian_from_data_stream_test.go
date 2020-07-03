package heap

import (
	"testing"
)

func TestMedianFinder_FindMedian(t *testing.T) {
	mf := Constructor()
	t.Logf("find=%f", mf.FindMedian())
	mf.AddNum(-1)
	t.Logf("find1=%f", mf.FindMedian())
	mf.AddNum(-2)
	t.Logf("find2=%f", mf.FindMedian())
	mf.AddNum(-3)
	t.Logf("find3=%f", mf.FindMedian())
	mf.AddNum(-4)
	t.Logf("find4=%f", mf.FindMedian())
	mf.AddNum(-5)
	t.Logf("find5=%f", mf.FindMedian())
}
