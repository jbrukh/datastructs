// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>.  All rights reserved.
// This software is governed by BSD-style license, see LICENSE file.

package hashset

import (
	"datastructs"
	. "datastructs/util"
	"testing"
)

type HashableInt int8

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
	New(10)
}

func TestPutContains(t *testing.T) {
	hs := New(100)
	for i := 0; i < 100; i++ {
		hs.Put(HashableInt(i))
		Assert(t, hs.Contains(HashableInt(i)), "doesn't contain %d", i)
	}
	Assert(t, !hs.Contains(HashableInt(100)), "contains something it shouldn't")
}

func BenchmarkPut(b *testing.B) {
	hs := New(2)
	BenchmarkFunc(b, func() {
		hs.Put(HashableInt(0))
	})
}
