package main

import (
	"fmt"
	"math"
)

func productExceptSelf(nums []int) []int {
	var extended []int

	for i := 0; i < len(nums); i++ {
		extended = append(extended, nums...)
	}

	fmt.Println(extended)

	var final []int
	acc := 1

	preVNum := nums[0]
	for i := 0; i < len(extended); i++ {

		numsIdx := int(math.Floor(float64(i / len(nums))))
		vNum := nums[numsIdx]

		fmt.Printf("numsIdx = %d, vNum = %d, prevVNum = %d, extended[i] = %d, acc = %d \n", numsIdx, vNum, preVNum, extended[i], acc)
		if preVNum != vNum {
			final = append(final, acc)
			acc = 1
			preVNum = vNum
		}

		if extended[i] != vNum {
			acc = acc * extended[i]
		}
	}

	final = append(final, acc)

	return final
}

func main() {
	v0 := []int{1, 2, 3, 4}
	productExceptSelf(v0)
	fmt.Printf("v0 = %d \n", v0)

	v1 := []int{-1, 1, 0, -3, 3}
	productExceptSelf(v1)
	fmt.Printf("v1 = %d \n", v1)

	v2 := []int{1, 2, 3, 4, 5}
	productExceptSelf(v2)
	fmt.Printf("v2 = %d \n", v2)

	v3 := []int{1, 0}
	productExceptSelf(v3)
	fmt.Printf("v3 = %d \n", v3)
}
