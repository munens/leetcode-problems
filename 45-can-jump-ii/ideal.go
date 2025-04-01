package main

import (
	"fmt"
	"math"
)

func jumpIdeal(nums []int) int {
	// set variables
	jumps := 0
	maxDistance := 0
	currentPos := 0

	// loop through all items, except the last.
	for i := 0; i < len(nums)-1; i++ {
		// look for what is greater, current reach or new reachable.
		maxDistance = int(math.Max(float64(maxDistance), float64(nums[i]+i)))

		fmt.Printf(
			"i = %d, nums[%d] = %d, maxDistance = %d, currentPos = %d, jumps = %d, cond: currentPos == i -> %v \n\n",
			i, i, nums[i], maxDistance, currentPos, jumps, currentPos == i,
		)

		// if this iteration is the current step
		if currentPos == i {

			// increment jumps, set current position to maxDistance.
			jumps++
			currentPos = maxDistance

			fmt.Printf("updated: jumps = %d, currentPos = %d \n\n", jumps, currentPos)
		}
	}

	// return total jumps
	return jumps
}

func main() {

	s := "min. number of jumps: %d \n\n"

	v0 := []int{2, 3, 1, 1, 4}
	fmt.Println("v0 = ", v0)
	fmt.Printf(s, jumpIdeal(v0))

	v1 := []int{2, 3, 0, 1, 4}
	fmt.Println("v2 => ", v1)
	fmt.Printf(s, jumpIdeal(v1))

	//v2 := []int{2, 5, 0, 0, 0}
	//fmt.Println(v2, " => ", v2)
	//fmt.Printf(s, jumpIdeal(v2))
	//
	//v3 := []int{0}
	//fmt.Println(v3, " => ", v3)
	//fmt.Printf(s, jumpIdeal(v3))
	//
	//v4 := []int{2, 0, 0}
	//fmt.Println(v4, " => ", v4)
	//fmt.Printf(s, jumpIdeal(v4))
	//
	//v5 := []int{2, 3, 4, 1, 4}
	//fmt.Println(v5, " => ", v5)
	//fmt.Printf(s, jumpIdeal(v5))
	//
	//v6 := []int{2, 3, 4, 1, 4, 0}
	//fmt.Println(v6, " => ", v6)
	//fmt.Printf(s, jumpIdeal(v6))
	//
	//v7 := []int{5, 0, 0, 0, 1, 0}
	//fmt.Println(v7, " => ", v7)
	//fmt.Printf(s, jumpIdeal(v7))
}
