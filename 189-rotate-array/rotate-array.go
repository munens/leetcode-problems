// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func rotate(nums []int, k int) {
	fmt.Println("nums = ", nums, ", k = ", k)
	thresh := len(nums) - k
	fmt.Println("threshI = ", thresh, ", threshV = ", nums[thresh])
	for i, _ := range nums {
		if thresh < len(nums) {
			currV := nums[i]
			threshV := nums[thresh]
			nums[i] = threshV
			nums[thresh] = currV
			thresh++
		} else {
			// handle array after thresh index
		}
	}

	fmt.Println(nums)
}

func main() {
	v0 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(v0, 3)
	fmt.Println("v0 = ", v0)
}
