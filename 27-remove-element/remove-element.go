// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"math"
)

func removeElement(nums []int, val int) int {
	k := 0

	i := 0
	finalIndex := len(nums) - 1

	for {
		fmt.Println("nums = ", nums, ", i = ", i, ", finalIndex = ", finalIndex, ", k = ", k)

		if i > len(nums)-1 {
			break
		}

		if nums[i] == val {

			nums[i] = nums[finalIndex]
			nums[finalIndex] = int(math.Inf(0))

			finalIndex--
			k++
		} else {
			i++
		}
	}

	return len(nums) - k
}

func main() {
	nums0 := []int{3, 2, 2, 3}
	v0 := removeElement(nums0, 3)
	fmt.Printf(" k = %d, nums = %d  \n\n", v0, nums0)

	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	v1 := removeElement(nums1, 2)
	fmt.Printf(" k = %d, nums = %d  \n\n", v1, nums1)
}
