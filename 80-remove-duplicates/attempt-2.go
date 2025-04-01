package main

import "fmt"

func removeDuplicatesAttempt2(nums []int, count int) int {
	fmt.Println("initial: ", nums)

	k := 0
	m := 1
	prevV := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		v := nums[i]

		if v == prevV {
			m++
		} else {
			m = 1
		}

		if m < count+1 {
			vi := nums[i]
			vj := nums[j]
			nums[i] = vj
			nums[j] = vi

			j++
		} else {
			k++
		}

		prevV = v
	}

	return len(nums) - k
}

func main() {
	v0 := []int{1, 1, 1, 2, 2, 3}
	k0 := removeDuplicatesAttempt2(v0, 2)
	fmt.Printf("final nums: %d, k = %d \n\n", v0, k0)

	v1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k1 := removeDuplicatesAttempt2(v1, 2)
	fmt.Printf("final nums: %d, k = %d \n\n", v1, k1)

	v2 := []int{0, 0, 0}
	k2 := removeDuplicatesAttempt2(v2, 2)
	fmt.Printf("final nums: %d, k = %d \n\n", v2, k2)

	v3 := []int{0, 0}
	k3 := removeDuplicatesAttempt2(v3, 2)
	fmt.Printf("final nums: %d, k = %d \n\n", v3, k3)

	v4 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 4}
	k4 := removeDuplicatesAttempt2(v4, 3)
	fmt.Printf("final nums: %d, k = %d \n\n", v4, k4)

	v5 := []int{0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4}
	k5 := removeDuplicatesAttempt2(v5, 2)
	fmt.Printf("final nums: %d, k = %d \n\n", v5, k5)
}
