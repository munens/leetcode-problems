package main

import "fmt"

func removeElementIdeal(nums []int, val int) int {
	fmt.Printf("nums = %d, val = %d \n\n", nums, val)
	j := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println("nums = ", nums, ", i = ", i, ", j = ", j)
		if nums[i] != val {
			v0 := nums[i]
			v1 := nums[j]

			nums[j] = v0
			nums[i] = v1
			//nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	return j
}

func main() {
	nums0 := []int{3, 2, 2, 3}
	v0 := removeElementIdeal(nums0, 3)
	fmt.Printf(" k = %d, nums = %d  \n\n", v0, nums0)

	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	v1 := removeElementIdeal(nums1, 2)
	fmt.Printf(" k = %d, nums = %d  \n\n", v1, nums1)
}
