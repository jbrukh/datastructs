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

// Locate will produce the location -- word and bit index --
// of the specified absolute index of the vector.
func locate(index uint) (word, bit uint) {
    return index / WORDSIZE, index % WORDSIZE
}

// max returns the maximum of two integers.
func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

// accomodate will grow the BitVector to accomodate new entries.
func (v *BitVector) accomodate(index uint) {
    words := int(index/WORDSIZE) + 1
    length := len(v.bits)
    if words > length { // reallocate
        newSlice := make([]byte, max(2*length, words))
        copy(newSlice, v.bits)
        v.bits = newSlice
    }
}

func (v *BitVector) accomodateLength(length int) {
    v.accomodate(uint(length*WORDSIZE) - 1)
}

// basisByte returns a byte with a single non-zero bit
// set at the specified index.
func basisByte(index uint) byte {
    // does NOT scale with WORDSIZE
    return 0x80 >> index
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

// Copy returns a copy of this BitVector.
func (v *BitVector) Copy() *BitVector {
    w := &BitVector{}
    w.accomodateLength(len(v.bits))
    copy(w.bits, v.bits)
    return w
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
    v.accomodateLength(length)
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
    v.accomodateLength(length)
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
