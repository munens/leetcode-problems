package main

import (
	"55-can-jump/tree"
	"fmt"
	"github.com/kr/pretty"
)

func main() {
	//v0 := []int{2, 3, 1, 1, 4}
	//resultV0 := tree.CanJump(v0)
	//fmt.Println(v0, " => ", resultV0)
	//pretty.Println("----------------------\n")
	//
	//v1 := []int{0}
	//resultV1 := tree.CanJump(v1)
	//fmt.Println(v1, " => ", resultV1)
	//pretty.Println("----------------------\n")
	//
	//v2 := []int{0, 1}
	//resultV2 := tree.CanJump(v2)
	//fmt.Println(v2, " => ", resultV2)
	//pretty.Println("----------------------\n")
	//
	//v3 := []int{2, 0, 0}
	//resultV3 := tree.CanJump(v3)
	//fmt.Println(v3, " => ", resultV3)
	//pretty.Println("----------------------\n")
	//
	//v4 := []int{3, 2, 1, 0, 4}
	//resultV4 := tree.CanJump(v4)
	//fmt.Println(v4, " => ", resultV4)
	//pretty.Println("----------------------\n")
	//
	v5 := []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	resultV5 := tree.NewTree(v5)
	fmt.Println(v5, " => ", resultV5)
	pretty.Println("----------------------\n")

	//CanJumpTree([]int{3, 2, 1, 0, 4})
	//
	//CanJumpTree([]int{3, 2, 1, 3, 4})
	//
	//CanJumpTree([]int{2, 5, 0, 0, 0})

}
