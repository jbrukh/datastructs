// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>.  All rights reserved.
// This software is governed by BSD-style license, see LICENSE file.

package datastructs

import (
	"crypto/md5"
	"gob"
)

var (
    digest = md5.New()
    encoder = gob.NewEncoder(digest)
)

// Hash creates an MD5 hash of an arbitrary object, and
// returns the value.  Serialization via gob is used to
// create the hash, and this function panics if the
// serialization cannot be performed.
func HashToBytes(obj interface{}) []byte {
    digest.Reset()
    if err := encoder.Encode(obj); err != nil {
        panic(err)
    }
    return digest.Sum()
}

// Returns a uint64 hash based on HashToBytes().
func Hash(obj interface{}) (code uint64) {
	hashBytes := HashToBytes(obj)
	for i := 0; i < 8; i++ {
		code |= (uint64(hashBytes[0]) << uint(i*8))
    }
	return
}

// Hashable provides an interface for hashable objects,
// analogous to Java's hashCode() and equals() methods.
type Hashable interface {
	HashCode() uint64
	Equal(other interface{}) bool
}
