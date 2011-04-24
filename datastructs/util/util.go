package util

import "testing"

// max returns the maximum of two integers.
func Max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

// min returns the maximum of two integers.
func Min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

// assert is the basic testing wrapper.
func Assert( t *testing.T, value bool, format string, args ...interface{} ) {
    if !value {
        t.Errorf(format,args...)
    }
}

func BenchmarkFunc(b *testing.B, someFunc func()) {
	for i := 0; i < b.N; i++ {
		someFunc()
    }
}
