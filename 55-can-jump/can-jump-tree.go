package main

import (
	"55-can-jump/tree"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func canJumpTreeFunc(nums []int, n *tree.Node, idx int) {
	finIdx := len(nums) - 1
	for _, conn := range n.Connectors {
		i := idx + conn.V
		
		var v int
		if i <= finIdx {
			v = nums[i]
		} else {
			v = nums[idx]
		}

		n := tree.NewNode(v, v != nums[finIdx])
		conn.End = n
		canJumpTreeFunc(nums, n, i)
	}

}

func CanJumpTree(nums []int) {

	fmt.Printf("nums = %d \n\n", nums)

	t := tree.NewTree(nums[0])
	canJumpTreeFunc(nums, t.Root, 0)
	spew.Dump(t)
}

func main() {
	CanJumpTree([]int{2, 3, 1, 1, 4})

	CanJumpTree([]int{3, 2, 1, 0, 4})

	CanJumpTree([]int{3, 2, 1, 3, 4})
	// 5/6: need to review how tree is built with this last array - which is unique as it the elements within it
	//      mean the 'jumping' can overshoot the array. In such a scenario we want to build the tree according to
	//      just what would be required to reach the last element in the array through the tree.
}
