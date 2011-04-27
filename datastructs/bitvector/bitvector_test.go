// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>.  All rights reserved.
// This software is governed by BSD-style license, see LICENSE file.

package bitvector

import (
	"testing"
	. "datastructs/util"
)

// Length of the BitVector
// for the benchmarking portion
const VECTOR_LEN = 1024

func TestAllocationSize(t *testing.T) {
	// test single byte allocation
	for i := 1; i <= 8; i++ {
		v := New(i)
		Assert(t, len(v.bits) == 1, "wrong allocation size for iteration %d", i)
	}
	// test double byte allocation
	for i := 9; i <= 16; i++ {
		v := New(i)
		Assert(t, len(v.bits) == 2, "wrong allocation size for iteration %d", i)
	}

	// test array growth
	v := New(1)
	v.Set(1, true)
	Assert(t, len(v.bits) == 1, "weird allocation after set")
	v.Set(7, true)
	Assert(t, len(v.bits) == 1, "weird allocation after set")
	v.Set(8, true)
	Assert(t, len(v.bits) == 2, "weird allocation after set (should have grown to 16)")
	w := New(1)
	w.Set(17, true)
	Assert(t, len(w.bits) == 3, "didn't grow correctly (should have grown to 3 bytes)")
}

func TestSetGet(t *testing.T) {
	v := New(8)
	w := New(16)
	v.Set(0, true)
	Assert(t, v.Get(0), "didn't set correctly")
	Assert(t, v.GetInt(0) == 1, "didn't set correctly")
	for i := 1; i < 8; i++ {
		Assert(t, !v.Get(i), "should be 0")
		Assert(t, v.GetInt(i) == 0, "GetInt should be 0")
	}

	// set all bits in w
	for i := 0; i < 16; i++ {
		w.Set(i, true)
	}
	// and check...
	for i := 0; i < 16; i++ {
		Assert(t, w.Get(i), "should be 1")
		Assert(t, w.GetInt(i) == 1, "GetInt should be 1")
	}
}

func TestNot(t *testing.T) {
	w := New(16)
	w.Not()
	// and check...
	for i := 0; i < 16; i++ {
		Assert(t, w.Get(i), "should be 1")
		Assert(t, w.GetInt(i) == 1, "GetInt should be 1")
	}
	w.Not()
	// and check...
	for i := 0; i < 16; i++ {
		Assert(t, !w.Get(i), "should be 0")
		Assert(t, w.GetInt(i) == 0, "GetInt should be 0")
	}
	w.Not()
	// and check...
	for i := 0; i < 16; i++ {
		Assert(t, w.Get(i), "should be 1")
		Assert(t, w.GetInt(i) == 1, "GetInt should be 1")
	}
}

func TestNotPackage(t *testing.T) {
	v := New(16)
	w := Not(v)
	Assert(t, v != w, "v and w are the same")
	// and check...
	for i := 0; i < 16; i++ {
		Assert(t, w.Get(i), "should be 1")
		Assert(t, w.GetInt(i) == 1, "GetInt should be 1")
	}
	z := Not(w)
	// and check...
	for i := 0; i < 16; i++ {
		Assert(t, !z.Get(i), "should be 0")
		Assert(t, z.GetInt(i) == 0, "GetInt should be 0")
	}
}

func TestCopy(t *testing.T) {
	v := New(16)
	w := v.Copy()
	Assert(t, v != w, "same object?")
	Assert(t, len(v.bits) == len(w.bits), "different byte allocations!")

}

func TestEqual(t *testing.T) {
	v := New(16)
	w := New(8)
	Assert(t, v.Equal(w), "%s should be equal to %s", v, w)
	Assert(t, w.Equal(v), "%s should be equal to %s", v, w)
	w.Not()
	Assert(t, !v.Equal(w), "%s should not be equal to %s", v, w)
	Assert(t, !w.Equal(v), "%s should not be equal to %s", v, w)
}

func TestOr(t *testing.T) {
	v, w, z := New(8), New(8), New(8)
	z.Not()
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			v.Set(i, true)
		} else {
			w.Set(i, true)
		}
	}
	v.Or(w)
	Assert(t, v.Equal(z), "v should all be ones: %s", v)
}

func TestAnd(t *testing.T) {
	v, w, z := New(8), New(8), New(8)
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			v.Set(i, true)
		} else {
			w.Set(i, true)
		}
	}
	v.And(w)
	Assert(t, v.Equal(z), "v should all be zeros: %s", v)
	z.Not()
	v.And(z)
	Assert(t, v.Equal(v), "v should be itself: %s", v)
	v.And(v)
	Assert(t, v.Equal(v), "v should be itself: %s", v)
}

func BenchmarkEqual(b *testing.B) {
	v, w := New(VECTOR_LEN), New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		v.Equal(w)
	}
}

func BenchmarkAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(VECTOR_LEN)
	}
}

func BenchmarkReallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(VECTOR_LEN)
		v.Set(VECTOR_LEN+100, true)
	}
}

func BenchmarkNot(b *testing.B) {
	v := New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		v.Not()
	}
}

func BenchmarkOr(b *testing.B) {
	v := New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		v.Or(v)
	}
}

func BenchmarkAnd(b *testing.B) {
	v := New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		v.And(v)
	}
}

func BenchmarkNotManual(b *testing.B) {
	v := New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		for j := 0; j < VECTOR_LEN; j++ {
			v.Set(j, true)
		}
	}
}

func BenchmarkSetTrue(b *testing.B) {
	v := New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		v.Set(0, true)
	}
}

func BenchmarkSetFalse(b *testing.B) {
	v := New(VECTOR_LEN)
	for i := 0; i < b.N; i++ {
		v.Set(0, false)
	}
}
