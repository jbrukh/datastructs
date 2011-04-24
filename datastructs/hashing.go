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
func Hash(obj interface{}) []byte {
    digest.Reset()
    if err := encoder.Encode(obj); err != nil {
        panic(err)
    }
    return digest.Sum()
}

// Hashable provides an interface for hashable objects,
// analogous to Java's hashCode() and equals() methods.
type Hashable interface {
	HashCode() int64
	Equal(other Hashable) bool
}
