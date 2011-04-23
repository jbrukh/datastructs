// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>
// All rights reserved.
// This software is governed by BSD-style license, see LICENSE.


package bitvector

import (
    "bytes"
    "fmt"
	. "datastructs/util"
)

// CONSTANTS
const WORDSIZE = 8


// BitVector represents a bit vector.
type BitVector struct {
    // internal representation (does NOT scale with WORDSIZE)
    bits []byte
}

// New creates a new BitVector that will accomodate at least
// the specified minimum length.  When necessary, the vector
// will automatically grow to accomodate the extra data. 
func New(length int) *BitVector {
    v := &BitVector{}
    v.accomodate(length)
    return v
}

// Length returns the current length of the BitVector in bits.
func (v *BitVector) Length() int {
    return WORDSIZE * len(v.bits)
}

// Clear will reset all bits to 0, but the size of the vector
// remains unaffected.
func (v *BitVector) Clear() {
    for inx, _ := range v.bits {
        v.bits[inx] = 0
    }
}

// Copy returns a copy of this BitVector.
func (v *BitVector) Copy() *BitVector {
    w := &BitVector{}
    w.accomodateBytes(len(v.bits))
    copy(w.bits, v.bits)
    return w
}

// String will create a string representation of the vector.
func (v *BitVector) String() string {
    buf := bytes.NewBufferString("")
    for _, word := range v.bits {
        for inx := 0; inx < WORDSIZE; inx++ {
            // if the inx-th bit is 1, result 
            // is 1; otherwise 0; following does NOT
            // scale with the value of WORDSIZE
            fmt.Fprint(buf, (0x80&word)>>7)
            word <<= 1
        }
    }
    return string(buf.Bytes())
}

// Equal returns true if and only if the BitVector
// provided is equal to this BitVector in the sense
// that all bit settings coincide in both structures.
func (v *BitVector) Equal(w *BitVector) bool {
	if w == nil {
		return false
	}
	minx := Max(len(v.bits),len(w.bits))
	for i := 0; i < minx; i++ {
		if v.getByte(i) != w.getByte(i) {
			return false
		}
	}
	return true
}

// Set will set the bit at the specified index to
// the specified boolean value.
func (v *BitVector) Set(index int, value bool) {
    v.accomodate(index+1)
    word, bitmask := locate(index)
    if value {
        v.bits[word] |= bitmask
    } else {
        v.bits[word] &= ^bitmask
    }
}

// IsSet returns true if and only if the bit at the
// specified index is set.
func (v *BitVector) Get(index int) bool {
    return v.GetInt(index) == 1
}

// Get will retrieve the value of the bit at the
// specified index as an signed integer.
func (v *BitVector) GetInt(index int) int {
    word, bitmask := locate(index)
    if word >= len(v.bits) {
        return 0
    }
    return int(v.bits[word] & bitmask >> uint(7 - index % WORDSIZE))
}

// Not negates this BitVector.
func (v *BitVector) Not() {
    for inx, _ := range v.bits {
        v.bits[inx] = ^v.bits[inx]
    }
}

// Not returns the negation of this BitVector.
func Not(v *BitVector) *BitVector {
    w := v.Copy()
    w.Not()
    return w
}

// Or will OR this BitVector with another one.
func (v *BitVector) Or(w *BitVector) {
	if w == nil {
		return
	}
    length := Max(len(v.bits), len(w.bits))
    for inx := 0; inx < length; inx++ {
		v.setByte(inx, v.getByte(inx) | w.getByte(inx))
	}
}

// Or returns the OR of two BitVectors.
func Or(v, w *BitVector) *BitVector {
    z := v.Copy()
    z.Or(w)
    return z
}

// And will AND this BitVector with another one.
func (v *BitVector) And(w *BitVector) {
    if w == nil {
		return
	}
	length := Max(len(v.bits), len(w.bits))
    for inx := 0; inx < length; inx++ {
		v.setByte(inx, v.getByte(inx) & w.getByte(inx))
	}
}

// And returns the AND of two BitVectors.
func And(v, w *BitVector) *BitVector {
    z := v.Copy()
    z.And(w)
    return z
}

// Locate will produce the location -- word and bit index --
// of the specified absolute index of the vector.
func locate(index int) (word int, bit byte) {
    return index / WORDSIZE, 0x80 >> uint(index % WORDSIZE)
}

// accomodate will grow the BitVector to accomodate new entries.
// The parameter must be positive.
func (v *BitVector) accomodate(elements int) {
    wordsNeeded := (elements - 1) / WORDSIZE + 1
    words := len(v.bits)
    if wordsNeeded > words { // reallocate
        newSlice := make([]byte, Max(2*words, wordsNeeded))
        copy(newSlice, v.bits)
        v.bits = newSlice
    }
}

// accomodate will grow the BitVector to accomodate the provided
// number of bytes.  The parameter must be positive.
func (v *BitVector) accomodateBytes(length int) {
    v.accomodate(length * WORDSIZE)
}

// getByte returns the value of the inx-th byte, if
// it is allocated, or 0.  That is, the vector is
// assumed to have an infinite 0-sequence extending
// to the right.
func (v *BitVector) getByte(inx int) byte {
    if inx < len(v.bits) {
        return v.bits[inx]
    }
    return byte(0)
}

// setByte will set the valie of the inx-th byte. If
// the interior byte array is not big enough, it may
// be reallocated if the data is non-trivial.
func (v *BitVector) setByte(inx int, data byte) {
    if inx >= len(v.bits) && data != 0 {
	    v.accomodateBytes(inx+1)	// will only resize for nontrivial data
	}
	v.bits[inx] = data
}
