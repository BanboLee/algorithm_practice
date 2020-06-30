package tree

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	root := obj.deserialize("[1,2,null,null,3,4,null,null,5,null,null]")
	data := obj.serialize(root)
	t.Logf(data)
}
