// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"math"
)

var MIN int = 0
var MAX int = int(math.Pow(10, 4))

//func maxProfit(prices []int) int {
//	var big = MIN
//	var small = MAX
//	indexOfBig := 0
//	indexOfSmall := 0
//	diff := 0
//	for i := 0; i < len(prices); i++ {
//		curr := prices[i]
//		if curr > big {
//			big = curr
//			indexOfBig = i
//		}
//
//		if curr < small {
//			small = curr
//			indexOfSmall = i
//		}
//
//		newDiff := big - small
//
//		fmt.Println("small = ", small, ", big = ", big, ", newDiff = ", newDiff, ", diff = ", diff)
//		if indexOfBig > indexOfSmall && newDiff > 0 && newDiff > diff {
//			diff = newDiff
//		}
//	}
//
//	return diff
//}

func maxProfit(prices []int) int {
	var small = MAX
	profit := 0
	for i := 0; i < len(prices); i++ {
		curr := prices[i]
		if curr < small {
			small = curr
		}

		newProfit := curr - small
		if newProfit > 0 && newProfit > profit {
			profit = newProfit
		}
	}

	return profit
}

func main() {
	p0 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Printf("p0 = %d \n\n", p0)

	p1 := maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Printf("p1 = %d \n\n", p1)

	p2 := maxProfit([]int{2, 4, 1})
	fmt.Printf("p2 = %d \n\n", p2)

}
