// Copyright (c) 2011.  Jake Brukhman <jbrukh@gmail.com>.  All rights reserved.
// This software is governed by BSD-style license, see LICENSE file.

package hashset

import (
	. "datastructs"
	"container/list"
)

// HashSet represents a hash set.
type HashSet struct {
	bins     map[uint64]*list.List
	capacity int
}

// New returns a new HashSet with the desired capacity. The capacity
// will have an impact of performance as it is directly proportional
// to collisions.
func New(capacity int) *HashSet {
	return &HashSet{
		capacity: capacity,
		bins:     make(map[uint64]*list.List, capacity),
	}
}

// listContains returns true if an only if the given list
// contains the given non-nil element.  This operation is
// O(n) at worst case.
func listContains(lst *list.List, obj Hashable) bool {
	if lst == nil {
		panic("you are providing a nil list")
	}
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
	hashCode := obj.HashCode() % uint64(this.capacity)
	lst, ok := this.bins[hashCode]
	if !ok {
		lst = list.New()
		this.bins[hashCode] = lst
	}
	return lst
}

// Put places the given element into the set. If the
// element is placed in for the first time this method
// will return true.
func (this *HashSet) Put(obj Hashable) bool {
	lst := this.getBin(obj)
	if listContains(lst, obj) {
		return false
	}
	lst.PushBack(obj)
	return true
}


func (this *HashSet) Remove(obj Hashable) (removed interface{}, ok bool) {
	lst := this.getBin(obj)
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
	hashCode := obj.HashCode() % uint64(this.capacity)
	if lst, ok := this.bins[hashCode]; ok {
		return listContains(lst, obj)
	}
	return false
}
