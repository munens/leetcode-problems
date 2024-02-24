package main

import (
	"fmt"
	"math"
)

const inf = math.MinInt

func shift(nums []int) {
	i := 0
	j := 0
	for {

		if i == len(nums) {
			break
		}

		fmt.Println("round: ", "nums = ", nums, ", i = ", i, ", j = ", j, ", vi = ", nums[i], ", vj = ", nums[j])

		if nums[i] != inf {
			vi := nums[i]
			vj := nums[j]
			nums[i] = vj
			nums[j] = vi
			j++
		}

		i++
	}
}

func removeDuplicates(nums []int) int {
	fmt.Println("initial nums:", nums)

	k := 0

	numsMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numsMap[v]++
		if numsMap[v] > 2 {
			nums[i] = inf
		} else {
			k++
		}
	}

	fmt.Println("prelim nums:", nums, numsMap)

	shift(nums)

	return k
}

func main() {
	v0 := []int{1, 1, 1, 2, 2, 3}
	k0 := removeDuplicates(v0)
	fmt.Printf("final nums: %d, k = %d \n\n", v0, k0)

	v1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k1 := removeDuplicates(v1)
	fmt.Printf("final nums: %d, k = %d \n\n", v1, k1)
}
