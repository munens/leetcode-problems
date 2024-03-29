// https://leetcode.com/problems/jump-game/submissions/1210880731/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

//func canJump(nums []int) bool {
//	for i := 0; i < len(nums); {
//		if i == len(nums)-1 {
//			return true
//		}
//
//		v := nums[i]
//		if v == 0 {
//			return false
//		}
//
//		i += v
//	}
//
//	return false
//}

//func canJump2(nums []int) bool {
//	acc := 0
//	for i := 0; i < len(nums); {
//		v := nums[i]
//		acc += v
//		if acc >= len(nums)-1 {
//			return true
//		} else if v == 0 {
//			return false
//		}
//
//		i += v
//	}
//
//	return false
//}

func canJumpFunc(nums []int, auxAmount int, currIdx int) bool {
	acc := 0
	for i := auxAmount; i < len(nums); {
		v := nums[i]
		acc += v
		if acc >= len(nums)-1 {
			return true
		}

		if nums[acc] == 0 && v > 0 {
			return canJumpFunc(nums, v-1, i)
		} else if v-1 == 0 {
			return false
		}

		i = acc
	}
}

func canJump3(nums []int) bool {

}

func main() {
	v0 := canJump2([]int{2, 3, 1, 1, 4})
	fmt.Printf("v0 = %t \n\n", v0)

	v1 := canJump2([]int{3, 2, 1, 0, 4})
	fmt.Printf("v1 = %t \n\n", v1)

	v2 := canJump2([]int{3, 2, 1, 3, 4})
	fmt.Printf("v2 = %t \n\n", v2)

	// read this: https://leetcode.com/problems/jump-game/description/comments/1728717
	// submitted the above 3/22/24 but was unable to understand why my solution was unable to beat all the test cases
	// based on what I understood about the question. The concept of jump above does not mean you have to jump 2 indexes
	// if the value at a specific index is 2 for example. It means you can jump 0, 1 or 2 spots.
	// - should create another function canJump2 to show progress in solving this solution.

	// My solution:
	// - managed to implement a way of going through array, nums and determine if it is possible to get to the last index
	//   so long as the values in the array are all positive.
	// - need to handle the case where an item in the array is 0. At this point we want to see the last value in the array
	//   that got us to where we are in the array and retry the last operation we did with a value, -1 less than we did
	//   last time - so long as the value for the new operation is not 0. If so then we conclude that it is impossible to reach
	//   the last index.

}
