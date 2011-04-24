package hashset

import (
	. "datastructs"
	"container/list"
)

type HashSet struct {
	bins map[uint64]*list.List
}

func New(capacity int) *HashSet {
	return &HashSet{
		make(map[uint64]*list.List, capacity),
	}
}

// Put places the given element into the set. If the
// element is placed in for the first time this method
// will return true.
func (this *HashSet) Put(obj Hashable) bool {
	hashCode := obj.HashCode()
	if lst, ok := this.bins[hashCode]; ok {
		for e := lst.Front(); e != nil; e = e.Next() {
			if obj.Equal(e.Value) {
				return false
			}
		}
		lst.PushBack(obj)
		return true
	}
	lst := list.New()
	lst.PushBack(obj)
	this.bins[hashCode] = lst
	return true
}

func (this *HashSet) Contains(obj Hashable) bool {
	hashCode := obj.HashCode()
	if lst, ok := this.bins[hashCode]; ok {
		for e := lst.Front(); e != nil; e = e.Next() {
			if obj.Equal(e.Value) {
				return true
			}
		}
	}
	return false
}
