// https://leetcode.com/problems/jump-game/submissions/1210880731/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"fmt"
	"github.com/kr/pretty"
)

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

//func canJumpFunc(nums []int, currIdx int, amount int) bool {
//	acc := 0
//	for i := currIdx; i < len(nums); {
//		n := nums[i] - amount
//		acc += n
//		if acc >= len(nums)-1 {
//			return true
//		}
//
//		if nums[acc] == 0 && n > 0 {
//			return canJumpFunc(nums, currIdx, n-1)
//		}
//
//		i = acc
//		amount = 0
//	}
//
//	return false
//}

//func canJump3(nums []int) bool {
//	return false
//}

func main() {
	//v0 := canJump3([]int{2, 3, 1, 1, 4})
	//fmt.Printf("v0 = %t \n\n", v0)
	//
	//v1 := canJump3([]int{3, 2, 1, 0, 4})
	//fmt.Printf("v1 = %t \n\n", v1)
	//
	//v2 := canJump3([]int{3, 2, 1, 3, 4})
	//fmt.Printf("v2 = %t \n\n", v2)

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
	//
	// My solution 4/23:
	// - After studying this problem I have decided the best way to solve this problem is to build a form of tree. A tree
	//   that allows us to traverse through the array considering all options and possibilities for jumps

	// My solution: 4/30
	// - take a look at my solution in ./tree folder. I am building a solution with Nodes and connectors. Whenever a node
	//   is created, a number of connectors are created for that node - depending on the value contained within a node. So if
	//   the node value is 5 then, there will be 5 connectors made with values 1 to 5 respectively.
	//   A connector has a 'jump value' which is the no. of jumps until the next value in the array. It has a Start value which
	//   is a reference to the previous node and an end Value that is a reference to the next node.
	//   plan: Create an algorithm to fill up the tree declared in that package with nodes and connectors until it it possible to
	//   create a tree for every array that can be provided.

	// 5/6: need to review how tree is built with this last array - which is unique as it the elements within it
	//      mean the 'jumping' can overshoot the array. In such a scenario we want to build the tree according to
	//      just what would be required to reach the last element in the array through the tree.
	//

	// 5/8: Tree model finally created that can tell us if it is possible to 'jump' beyond the end of the array. The
	//      tree does this by creating a leaf node to indicate that that branch has no children. This leaf node has a
	//      value of -1. Next steps - to answer the question of the exercise - need to find a way to traverse the
	//      entire tree and find a value of -1. If that can be found that means that canJump can equal true, otherwise
	//      false. Could also implement tree with linkedlist representing connectors on a node. Currently this is
	//      implemented with arrays.

	v0 := []int{2, 3, 1, 1, 4}
	resultV0 := canJumpIdeal(v0)
	fmt.Println(v0, " => ", resultV0)
	pretty.Println("----------------------\n")

	v1 := []int{0}
	resultV1 := canJumpIdeal(v1)
	fmt.Println(v1, " => ", resultV1)
	pretty.Println("----------------------\n")

	v2 := []int{0, 1}
	resultV2 := canJumpIdeal(v2)
	fmt.Println(v2, " => ", resultV2)
	pretty.Println("----------------------\n")

	v3 := []int{2, 0, 0}
	resultV3 := canJumpIdeal(v3)
	fmt.Println(v3, " => ", resultV3)
	pretty.Println("----------------------\n")

	v4 := []int{3, 2, 1, 0, 4}
	resultV4 := canJumpIdeal(v4)
	fmt.Println(v4, " => ", resultV4)
	pretty.Println("----------------------\n")

	v5 := []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	resultV5 := canJumpIdeal(v5)
	fmt.Println(v5, " => ", resultV5)
	pretty.Println("----------------------\n")
}
