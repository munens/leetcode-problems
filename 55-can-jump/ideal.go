package main

import (
	"fmt"
	"github.com/kr/pretty"
)

func canJumpIdeal(nums []int) bool {
	farthest := 0

	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}
		farthest = maxFarthest(farthest, i+nums[i])
		fmt.Println("farthest = ", farthest)
	}

	return farthest >= len(nums)-1
}

func maxFarthest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//v0 := []int{2, 3, 1, 1, 4}
	//resultV0 := canJumpIdeal(v0)
	//fmt.Println(v0, " => ", resultV0)
	//pretty.Println("----------------------\n")
	//
	//v1 := []int{0}
	//resultV1 := canJumpIdeal(v1)
	//fmt.Println(v1, " => ", resultV1)
	//pretty.Println("----------------------\n")
	//
	//v2 := []int{0, 1}
	//resultV2 := canJumpIdeal(v2)
	//fmt.Println(v2, " => ", resultV2)
	//pretty.Println("----------------------\n")
	//
	//v3 := []int{2, 0, 0}
	//resultV3 := canJumpIdeal(v3)
	//fmt.Println(v3, " => ", resultV3)
	//pretty.Println("----------------------\n")
	//
	//v4 := []int{3, 2, 1, 0, 4}
	//resultV4 := canJumpIdeal(v4)
	//fmt.Println(v4, " => ", resultV4)
	//pretty.Println("----------------------\n")
	//
	//v5 := []int{2, 5, 0, 0, 0}
	//resultV5 := canJumpIdeal(v5)
	//fmt.Println(v5, " => ", resultV5)
	//pretty.Println("----------------------\n")

	v6 := []int{1, 0, 5, 0, 0}
	resultV6 := canJumpIdeal(v6)
	fmt.Println(v6, " => ", resultV6)
	pretty.Println("----------------------\n")

	//v7 := []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	//resultV7 := canJumpIdeal(v7)
	//fmt.Println(v7, " => ", resultV7)
	//pretty.Println("----------------------\n")
}
