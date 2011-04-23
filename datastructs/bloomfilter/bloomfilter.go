package bloomfilter

import (
	"crypto/md5"
	"gob"
	. "datastructs/bitvector"
    . "datastructs/util"
)

var (
    digest = md5.New()
    encoder = gob.NewEncoder(digest)
	x = Max(1,2)
)

type BloomFilter struct {
	v *BitVector
}

// Hash creates an MD5 hash of an arbitrary object, and
// returns the value.  Serialization via gob is used to
// create the hash, and this function panics if the
// serialization cannot be performed.
func Hash(obj interface{}) []byte {
    digest.Reset() 
    if err := encoder.Encode(obj); err != nil {
        panic(err)
    }
    return digest.Sum()
}
