package main

import "fmt"

func rotateIdeal(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		vs := nums[start]
		ve := nums[end]
		nums[start] = ve
		nums[end] = vs
		start++
		end--
	}
}

func main() {
	v0 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateIdeal(v0, 3)
	fmt.Printf("v0 = %d \n\n", v0)

	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotateIdeal(v1, 4)
	fmt.Printf("v1 = %d \n\n", v1)

	v2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotateIdeal(v2, 199)
	fmt.Printf("v2 = %d \n\n", v2)
}
