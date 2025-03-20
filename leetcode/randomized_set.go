package leetcode

import (
	"maps"
	"math/rand"
	"slices"
)

/**
 * 380. Insert Delete GetRandom O(1)
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// It looks like most of this should use a HashMap (or map)
// Problem: Need some way to keep track of what's valid in O(1)
type RandomizedSet struct {
	set map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]bool)}
}

func (this *RandomizedSet) Insert(val int) bool {
	initVal := this.set[val]
	this.set[val] = true

	return !initVal
}

func (this *RandomizedSet) Remove(val int) bool {
	initVal := this.set[val]
	delete(this.set, val)

	return initVal
}

func (this *RandomizedSet) GetRandom() int {
	// slices.Collect is the same as the collect fn in Rust
	// map.Keys returns an iter
	keys := slices.Collect(maps.Keys(this.set))

	return keys[rand.Intn(len(keys))]
}
