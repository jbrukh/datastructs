// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>
// All rights reserved.
// This software is governed by BSD-style license, see LICENSE.


package bitvector

import (
    "bytes"
    "fmt"
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
func New(length uint) *BitVector {
    v := &BitVector{}
    v.accomodate(length - 1)
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
	minx := max(len(v.bits),len(w.bits))
	for i := 0; i < minx; i++ {
		if v.getByte(i) != w.getByte(i) {
			return false
		}
	}
	return true
}

// Set will set the bit at the specified index to
// the specified boolean value.
func (v *BitVector) Set(index uint, value bool) {
    v.accomodate(index)
    word, bit := locate(index)
    if value {
        v.bits[word] |= basisByte(bit)
    } else {
        v.bits[word] &= ^basisByte(bit)
    }
}

// IsSet returns true if and only if the bit at the
// specified index is set.
func (v *BitVector) Get(index uint) bool {
    return v.GetInt(index) == 1
}

// Get will retrieve the value of the bit at the
// specified index as an signed integer.
func (v *BitVector) GetInt(index uint) int {
    word, bit := locate(index)
    if int(word) >= len(v.bits) {
        return 0
    }
    return int((v.bits[word] & basisByte(bit)) >> (7 - bit))
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
    length := len(w.bits)
    v.accomodateBytes(length)
    for inx, _ := range v.bits {
        if inx < length {
            v.bits[inx] |= w.bits[inx]
        }
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
    length := len(w.bits)
    v.accomodateBytes(length)
    for inx, _ := range v.bits {
        if inx < length {
            v.bits[inx] &= w.bits[inx]
        } else {
            v.bits[inx] = 0 //
        }
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
func locate(index uint) (word, bit uint) {
    return index / WORDSIZE, index % WORDSIZE
}

// accomodate will grow the BitVector to accomodate new entries.
func (v *BitVector) accomodate(index uint) {
    if index < 0 {
        panic("must specify nonnegative index")
    }
    words := int(index/WORDSIZE) + 1
    length := len(v.bits)
    if words > length { // reallocate
        newSlice := make([]byte, max(2*length, words))
        copy(newSlice, v.bits)
        v.bits = newSlice
    }
}

func (v *BitVector) accomodateBytes(length int) {
    v.accomodate(uint(length*WORDSIZE) - 1)
}

// basisByte returns a byte with a single non-zero bit
// set at the specified index.
func basisByte(index uint) byte {
    // does NOT scale with WORDSIZE
    return 0x80 >> index
}

// getByte returns the value of the inx-th byte, if
// it is allocated, or 0.  That is, the vector is
// assumed to have an infinite 0-sequence extending
// to the right.
func (v *BitVector) getByte(inx int) byte {
    if inx < len(v.bits) {
        return v.bits[inx]
    }
    return 0
}

// max returns the maximum of two integers.
func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

// min returns the maximum of two integers.
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
