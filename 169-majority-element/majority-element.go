// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func majorityElement(nums []int) int {
	numsMap := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			numsMap[v]++
		} else {
			numsMap[v] = 1
		}
	}

	fmt.Println(numsMap)
	maxV := 0
	elem := 0
	for i, v := range numsMap {
		if val, _ := numsMap[v]; val > maxV {
			maxV = numsMap[v]
			elem = i
		}
	}

	return elem
}

func main() {
	v0 := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println("v0 = ", v0)

	v1 := majorityElement([]int{2, 2, 1, 1, 1, 2, 2, 3, 3, 3, 6, 7, 7, 7, 7, 7})
	fmt.Println("v1 = ", v1)

	v3 := majorityElement([]int{6, 5, 5})
	fmt.Println("v3 = ", v3)
}
