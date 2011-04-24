package bloomfilter

import (
	. "datastructs/bitvector"
    . "datastructs/util"
)

type BloomFilter struct {
	v *BitVector
}
