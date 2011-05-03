// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>.  All rights reserved.
// This software is governed by BSD-style license, see LICENSE file.

package prefixtree

import (
	. "datastructs/util"
	"testing"
)

func TestContainsEmpty(t *testing.T) {
	pt := New()
	Assert(t, pt.Contains(""), "where's the empty word?")
}

func TestPutContains(t *testing.T) {
	pt := New()
	pt.Put("a")
	Assert(t, pt.Contains("a"), "a!")
	pt.Put("ab")
	Assert(t, pt.Contains("ab"), "ab!")
    Assert(t, !pt.Contains("b"), "b!!")
	pt.Put("abel")
	Assert(t, pt.Contains("abel"), "abel!")
	Assert(t, !pt.Contains("abe"), "abel!")
}

func TestGetChild(t *testing.T) {
	root := newNode()
	n1, n2, n3 := getChild(root,'a'), getChild(root,'a'), getChild(root,'b')
	Assert(t, n1==n2, "faulty getChild")
	Assert(t, n1!=n3, "faulty getChild")
}
