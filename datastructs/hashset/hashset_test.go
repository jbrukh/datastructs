package hashset

import (
	"datastructs"
	"testing"
	"fmt"
)

type HashableInt int8;

func (this HashableInt) HashCode() uint64 {
	return datastructs.Hash(this)
}

func (this HashableInt) Equal(obj interface{}) bool {
	if t, ok := obj.(HashableInt); ok && t == this {
		return true
	}
	return false
}

func TestNew(t *testing.T) {
	hs := New(10)
	hs.Put( HashableInt(11) )
	fmt.Println(hs.Contains( HashableInt(11) ))
}
