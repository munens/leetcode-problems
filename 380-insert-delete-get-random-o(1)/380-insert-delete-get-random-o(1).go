// https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
)

type RandomizedSet struct {
	Size   int
	Keys   []int
	Values map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Size:   0,
		Keys:   []int{},
		Values: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.Values[val]; ok {
		return false
	}

	rs.Keys = append(rs.Keys, val)
	rs.Values[val] = len(rs.Keys) - 1
	rs.Size += 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if currIdx, ok := rs.Values[val]; ok {
		lastIdx := len(rs.Keys) - 1
		rs.Keys[currIdx] = rs.Keys[lastIdx]
		rs.Values[rs.Keys[lastIdx]] = currIdx

		rs.Keys = rs.Keys[:lastIdx]
		delete(rs.Values, val)
		rs.Size -= 1

		return true
	} else {
		return false
	}
}

func (rs *RandomizedSet) Traverse() {
	for k, v := range rs.Values {
		fmt.Printf("key: %d, Value: %d \n\n", k, v)
	}
}

func genRandomNumberBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (rs *RandomizedSet) GetRandom() int {
	randomIdx := genRandomNumberBetween(0, len(rs.Keys)-1)
	fmt.Printf("random v = %d \n\n", randomIdx)
	return rs.Keys[randomIdx]
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

	rs.Insert(3)
	spew.Dump(rs)

	rs.Insert(4)
	spew.Dump(rs)

	rs.Remove(2)
	spew.Dump(rs)

	rs.Remove(3)
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
//   with its index in Keys on the RandomizedSet (i.e. Value{ V int Idx int}). The problem is that if an item is
//   deleted, this index can also be deleted. But it also means every other value on the RandomizedSet also has to have
//   its index in the Keys also updated. The indexes would be out of sync.
//
