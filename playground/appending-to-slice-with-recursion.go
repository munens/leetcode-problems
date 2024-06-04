package main

import (
	"fmt"
)

func main() {
	//s := []int{}
	//add(s)
	//fmt.Printf("%#v\n", s)
	//
	//s = []int{}
	//s = addReturn(s)
	//fmt.Printf("%#v\n", s)
	//
	//s = []int{}
	//r := &s
	//addReference(r)
	//fmt.Printf("%#v\n", s)

	v := 1
	t := []*int{&v}
	addReference2(&t)

	fmt.Printf("%#v\n", t)
}

// whatever terminates the recursion
func finalCondition(s []int) bool {
	if len(s) >= 10 {
		return true
	}
	return false
}

// incorrect
func add(s []int) {
	if finalCondition(s) {
		return
	}
	s = append(s, 1)
	add(s)
}

// correct way to return from recursion
func addReturn(s []int) []int {
	if finalCondition(s) {
		return s
	}
	s = append(s, 1)
	return addReturn(s)
}

// using a reference
func addReference(s *[]int) {
	if finalCondition(*s) {
		return
	}
	*s = append(*s, 1)
	addReference(s)
}

// whatever terminates the recursion
func finalCondition2(s []*int) bool {
	if len(s) >= 10 {
		return true
	}
	return false
}

// using a reference
func addReference2(s *[]*int) {
	fmt.Printf("%#v\n", s)
	if finalCondition2(*s) {
		return
	}
	v := 1
	*s = append(*s, &v)
	addReference2(s)
}
