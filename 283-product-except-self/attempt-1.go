package main

import (
	"fmt"
	"math"
)

func productExceptSelfAttempt1(nums []int) []int {
	var extended []int

	for i := 0; i < len(nums); i++ {
		extended = append(extended, nums...)
	}

	fmt.Println(extended)

	var final []int
	acc := 1
	prevNumsIdx := 0

	for i := 0; i < len(extended); i++ {

		numsIdx := int(math.Floor(float64(i / len(nums))))
		vNum := nums[numsIdx]

		fmt.Printf("numsIdx = %d, vNum = %d, prevNumsIdx = %d, extended[i] = %d, acc = %d \n", numsIdx, vNum, prevNumsIdx, extended[i], acc)
		if prevNumsIdx != numsIdx {
			final = append(final, acc)
			acc = 1
			prevNumsIdx = numsIdx
		}

		if extended[i] != vNum {
			acc = acc * extended[i]
		}
	}

	if acc == 1 {
		final = nums
	} else {
		final = append(final, acc)
	}

	return final
}

func main() {
	v0 := []int{1, 2, 3, 4}
	//f0 := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	vFin0 := productExceptSelfAttempt1(v0)
	fmt.Printf("v0 = %d \n", vFin0)

	v1 := []int{-1, 1, 0, -3, 3}
	vFin1 := productExceptSelfAttempt1(v1)
	fmt.Printf("v1 = %d \n", vFin1)

	v2 := []int{1, 2, 3, 4, 5}
	vFin2 := productExceptSelfAttempt1(v2)
	fmt.Printf("v2 = %d \n", vFin2)

	v3 := []int{1, 0}
	//f3 := []int{1, 0, 1, 0}
	vFin3 := productExceptSelfAttempt1(v3)
	fmt.Printf("v3 = %d \n", vFin3)

	v4 := []int{0, 0}
	vFin4 := productExceptSelfAttempt1(v4)
	fmt.Printf("v4 = %d \n", vFin4)

	v5 := []int{2, 2, 2}
	//f5 := []int{2, 2, 2, 2, 2, 2, 2, 2, 2}
	vFin5 := productExceptSelfAttempt1(v5)
	fmt.Printf("v5 = %d \n", vFin5)
}
