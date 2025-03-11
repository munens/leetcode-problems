// https://leetcode.com/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
)

/* 1st attempt
func jump(nums []int) int {
	var thresh = 0
	count := 0

	if len(nums) == 1 {
		return count
	}

	for i := 0; i < len(nums); i++ {
		newThresh := maxV(thresh, i+nums[i])
		if newThresh > thresh {
			thresh = newThresh
			count += 1
		}

		if thresh >= len(nums)-1 {
			return count
		}
	}

	return count
}

func maxV(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// solution 9/10/24
// perhaps there is a way to do a similar traversal to the ideal solution in 55-can-jump. But while doing that chart
// all the possible ways we can jump through the nums array.
// other proposal - perhaps we can augment a count value any time the threshold value is greater than the previous one
// - if the threshold goes up, that means that we have done a progressive legitimate jump in the nums array (and we want
//   to count that).
//
// 9/11-13/24
// - attempted a solution like the one mentioned above. Did not work

*/

func getJumpPossibilities(nums []int) [][]int {
	var poss [][]int
	for i := 0; i < len(nums); i++ {
		var jumps []int

		if nums[i] == 0 {
			jumps = append(jumps, 0)
		}

		for j := 1; j <= nums[i]; j++ {
			jumps = append(jumps, j)
		}

		poss = append(poss, jumps)
	}

	return poss
}

// arr = [2, 3, 1, 1, 4]
// curr = [[1, 2], [1, 2, 3], [1], [1], [1, 2, 3, 4]]

// arr = [2, 3, 0, 1, 4]
// curr = [[1, 2], [1, 2, 3], [0], [1], [1, 2, 3, 4]]

// currObj = {
// 	0:
// }

func getJumps(res *[][]int, poss [][]int, inner []int, acc []int, outerIdx int) {
	possLen := len(poss)
	for i := 0; i < len(inner); i++ {
		v := inner[i]
		fmt.Println("i = ", i, ", v = ", v, ", outerIdx = ", outerIdx, ", inner = ", inner, ", acc = ", acc)
		var newOuterIdx int
		if v == 0 {
			break
		} else {
			newOuterIdx = v + outerIdx
		}
		fmt.Println("newOuterIdx = ", newOuterIdx)

		if newOuterIdx >= possLen {
			if newOuterIdx == possLen {
				*res = append(*res, acc)
				fmt.Println(res)
				fmt.Println("--------------------")
			} else {
				break
			}
		} else {
			getJumps(res, poss, poss[newOuterIdx], append(acc, v), newOuterIdx)
		}
	}
}

func jump(nums []int) {
	poss := getJumpPossibilities(nums)

	fmt.Println("poss => ", poss)
	var jumps [][]int
	getJumps(&jumps, poss, poss[0], []int{}, 0)
	fmt.Println("jumps => ", jumps)
}

func main() {
	v0 := []int{2, 3, 1, 1, 4}
	fmt.Println("v0 = ", v0)
	jump(v0)

	fmt.Println("\n")

	v1 := []int{2, 3, 0, 1, 4}
	fmt.Println(v1, " => ", v1)
	jump(v1)

	fmt.Println("\n")

	v2 := []int{2, 5, 0, 0, 0}
	fmt.Println(v2, " => ", v2)
	jump(v2)

	//
	//v3 := []int{0}
	//resV3 := jump(v3)
	//fmt.Println(v3, " => ", resV3)
	//
	//v4 := []int{2, 0, 0}
	//resV4 := jump(v4)
	//fmt.Println(v4, " => ", resV4)
	//
	//v6 := []int{2, 3, 4, 1, 4}
	//resV6 := jump(v6)
	//fmt.Println(v6, " => ", resV6)
	//
	//v7 := []int{2, 3, 4, 1, 4, 0}
	//resV7 := jump(v7)
	//fmt.Println(v7, " => ", resV7)
	//
	//v5 := []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	//resV5 := jump(v5)
	//fmt.Println(v5, " => ", resV5)

	// 9/16
	// - attempting new approach will attempt to get all the possibilities in terms of jumps based on nums
	//   e.g. for nums = []int{2, 3, 1, 1, 4} all the possibilities are: [[1, 2], [1, 2, 3], [1], [1], [1, 2, 3, 4]]
	//   now will attempt to get the minimum amount of jumps based on knowing all the possibilities above e.g. 1, 2, 1. 3

	// 10/7
	// - gotten all possibilities function to work properly. Trying to find out why jumps function is ending primeturely.
	//   Seems like not all possible arrays re being added to 'res'.
}
