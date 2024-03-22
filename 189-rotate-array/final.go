package main

import "fmt"

func rotateFinal(nums []int, k int) {
	rotateIndex := k % len(nums)
	threshIndex := len(nums) - rotateIndex
	indexed := make([]int, len(nums))
	final := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i < rotateIndex {
			indexed[i] = threshIndex
			threshIndex++
		} else {
			indexed[i] = i - rotateIndex
		}

		final[i] = nums[indexed[i]]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = final[i]
	}
}

func main() {
	v0 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateFinal(v0, 3)
	fmt.Printf("v0 = %d \n\n", v0)

	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotateFinal(v1, 4)
	fmt.Printf("v1 = %d \n\n", v1)

	v2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotateFinal(v2, 199)
	fmt.Printf("v2 = %d \n\n", v2)
}
