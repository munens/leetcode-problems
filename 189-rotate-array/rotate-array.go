// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func rotate(nums []int, k int) {

	if len(nums) == 1 || k == 0 {
		return
	}
	fmt.Println("nums = ", nums, ", k = ", k)
	initThresh := len(nums) - k
	thresh := initThresh
	fmt.Println("threshI = ", thresh, ", threshV = ", nums[thresh])
	for i, _ := range nums {
		if thresh < len(nums) {
			currV := nums[i]
			threshV := nums[thresh]
			nums[i] = threshV
			nums[thresh] = currV
			thresh++
		} else {
			break
		}
	}

	fmt.Println("init nums: ", nums)

	i := k
	j := initThresh
	finIndex := len(nums) - 1

	for {
		if i > finIndex {
			break
		}

		if j > finIndex {
			j = finIndex
		}

		vi := nums[i]
		vj := nums[j]

		nums[i] = vj
		nums[j] = vi

		i++
		j++
	}

	fmt.Println("final nums: ", nums)
}

func rotate2(nums []int, k int) {
	fmt.Println("nums = ", nums, ", k = ", k)
	rotateIndex := k % len(nums)
	threshIndex := len(nums) - rotateIndex
	indexed := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i < rotateIndex {
			indexed[i] = threshIndex
			threshIndex++
		} else {
			indexed[i] = i - rotateIndex
		}
	}

	fmt.Println("indexed: ", indexed)
}

func rotate3(nums []int, k int) {
	fmt.Println("nums = ", nums, ", k = ", k)
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
	//v0 := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate(v0, 3)
	//fmt.Printf("v0 = %d \n\n", v0)
	//
	//v1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//rotate(v1, 9)
	//fmt.Printf("v1 = %d \n\n", v1)

	// 2/27: some cases to handle
	// - what happens if k > len(nums) ? How to handle this use case
	// - Need to go through algorithm right after line 27, does it work for all cases? It appears not to

	//v0 := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate2(v0, 3)
	//fmt.Printf("v0 = %d \n\n", v0)
	//
	//v1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//rotate2(v1, 4)
	//fmt.Printf("v1 = %d \n\n", v1)
	//
	//v2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//rotate2(v2, 199)
	//fmt.Printf("v2 = %d \n\n", v2)

	// 3/1: look at rotate2
	// - I have solved the problem of k > len(nums). We use modulo to find the remainder of the division of k and len(nums). This gives
	//   a different value of k that then allows us to know where to start rotating the array.
	// - I have created an indexed array above. It shows the correct indexes for the array that is supposed to be indexed.
	// - need to use the indexed array to set the nums array value in the right place, and this should be easy 2*O(n)

	v0 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(v0, 3)
	fmt.Printf("v0 = %d \n\n", v0)

	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate3(v1, 4)
	fmt.Printf("v1 = %d \n\n", v1)

	v2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate3(v2, 199)
	fmt.Printf("v2 = %d \n\n", v2)

}
