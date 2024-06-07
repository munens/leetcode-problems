package main

import (
	"55-can-jump/tree"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func CanJumpTree(nums []int) {

	fmt.Printf("nums = %d \n\n", nums)

	t := tree.NewTree(nums)

	spew.Dump(t)

	//t.RoundTrip()
}

func main() {
	CanJumpTree([]int{2, 3, 1, 1, 4})

	//CanJumpTree([]int{3, 2, 1, 0, 4})
	//
	//CanJumpTree([]int{3, 2, 1, 3, 4})
	//
	//CanJumpTree([]int{2, 5, 0, 0, 0})
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
}
