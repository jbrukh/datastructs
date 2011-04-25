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

// listContains returns true if an only if the given list
// contains the given non-nil element.  This operation is
// O(n) at worst case.
func listContains(lst *list.List, obj Hashable) bool {
	for e := lst.Front(); e != nil; e = e.Next() {
        if obj.Equal(e.Value) {
			return true
        }
	}
	return false
}

// getBin returns the bin associated with the hash of the given
// object.  If the bin does not exist, it is created and placed
// in the internal map.
func (this *HashSet) getBin(obj Hashable) *list.List {
	hashCode := obj.HashCode()
	var lst *list.List
	if _, ok := this.bins[hashCode]; !ok {
		lst := list.New()
		this.bins[hashCode] = lst
	}
	return lst
}

// Put places the given element into the set. If the
// element is placed in for the first time this method
// will return true.
func (this *HashSet) Put(obj Hashable) bool {
	hashCode := obj.HashCode()
	lst := getBin(obj)
	if listContains(lst, obj) {
		return false
	}
	lst.PushBack(obj)
	return true
}


func (this *HashSet) Remove(obj Hashable) (interface{}, ok bool) {
	hashCode := obj.HashCode()
	lst := getBin(obj)
	for e := lst.Front(); e != nil; e = e.Next() {
        if obj.Equal(e.Value) {
            return lst.Remove(e), true
        }
    }
	return nil, false
}

// Contains returns true if and only if this set contains
// the given object.
func (this *HashSet) Contains(obj Hashable) bool {
	hashCode := obj.HashCode()
    if lst, ok := this.bins[hashCode]; ok {
		return listContains(lst,obj)
	}
	return false
}
