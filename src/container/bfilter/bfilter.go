package bfilter

import (
    . "container/bitvector"
    "crypto/md5"
    "encoding/binary"
    //"fmt"
)


type BloomFilter struct {
    v *BitVector
}

func Hash( obj interface{} ) []byte {
    digest := md5.New()
    if err := binary.Write(digest, binary.LittleEndian, obj); err != nil {
        panic(err)
    }
    result := digest.Sum()
    return result
}

