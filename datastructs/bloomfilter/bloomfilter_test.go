package bloomfilter

import (
	. "datastructs/util"
	"testing"
)

type testingStruct struct {
	I int
    j int64
    k complex128
}

func TestTrivial(t *testing.T) {

}

// BenchmarkHashTime tests how long it takes to hash the testingStruct.
func BenchmarkHashTime(b *testing.B) {
	BenchmarkFunc(b,func() { Hash(testingStruct{}) })
}
