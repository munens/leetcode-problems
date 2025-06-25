// https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
)

type Value struct {
	V   int
	Idx int
}

type RandomizedSet struct {
	Size        int
	KeysInOrder []int
	Values      map[int]Value
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Size:        0,
		KeysInOrder: []int{},
		Values:      make(map[int]Value),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.Values[val]; ok {
		return false
	}

	rs.KeysInOrder = append(rs.KeysInOrder, val)
	rs.Values[val] = Value{V: val, Idx: len(rs.KeysInOrder) - 1}
	rs.Size += 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if v, ok := rs.Values[val]; ok {
		idx := v.Idx
		rs.KeysInOrder = append(rs.KeysInOrder[:idx], rs.KeysInOrder[idx+1:]...)
		delete(rs.Values, val)
		rs.Size -= 1
		return true
	} else {
		return false
	}
}

func genRandomNumberBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (rs *RandomizedSet) GetRandom() int {
	randomIdx := genRandomNumberBetween(0, len(rs.KeysInOrder)-1)
	fmt.Printf("random v = %d \n\n", randomIdx)
	return rs.KeysInOrder[randomIdx]
}

func main() {
	rs := Constructor()
	rs.Insert(0)
	spew.Dump(rs)
	rs.Insert(1)
	spew.Dump(rs)
	rs.Remove(0)
	spew.Dump(rs)

	rs.Insert(2)
	spew.Dump(rs)

	rs.Remove(1)
	spew.Dump(rs)

	v := rs.GetRandom()
	fmt.Printf("random1: %d \n\n", v)

	v1 := rs.GetRandom()
	fmt.Printf("random2: %d \n\n", v1)

	v2 := rs.GetRandom()
	fmt.Printf("random3: %d \n\n", v2)
}

// 6/24:
// - The concern with this strategy here is generating a random number to pull from the RandomizedSet. It means we have
//   know what keys are available on th RandomizedSet at all times. This presents the issue of deletion of items from
//   the RandomizedSet. If we delete a value from the set it means we have to also delete a value from our store
//   of all the keys available in the RandomizedSet. I tried to solve this problem initially by storing a value together
//   with its index in KeysInOrder on the RandomizedSet (i.e. Value{ V int Idx int}). The problem is that if an item is
//   deleted, this index can also be deleted. But it also means every other value on the RandomizedSet also has to have
//   its index in the KeysInOrder also updated. The indexes would be out of sync.
//
