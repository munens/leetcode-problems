// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

// first solution
func maxProfit(prices []int) int {
	fmt.Printf("prices = %d \n\n", prices)
	profit := 0
	counter := 0
	for i := 1; i < len(prices); i++ {
		curr := prices[i]
		prev := prices[i-1]

		fmt.Println("prev = ", prev, ", curr = ", curr)

		if curr > prev {
			counter += curr - prev

			fmt.Println("counter = ", counter)

			if i == len(prices)-1 {
				profit += counter
			}
		}

		if curr <= prev {
			profit += counter
			counter = 0
		}

		fmt.Printf("latest profit = %d \n\n", profit)
	}

	return profit
}

// second solution
func maxProfit2(prices []int) int {
	fmt.Printf("prices = %d \n\n", prices)
	profit := 0
	for i := 1; i < len(prices); i++ {
		curr := prices[i]
		prev := prices[i-1]

		fmt.Println("prev = ", prev, ", curr = ", curr)

		if curr > prev {
			profit += curr - prev
			fmt.Printf("latest profit = %d \n\n", profit)
		}
	}

	return profit
}

func main() {
	p0 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Printf("p0 = %d \n\n", p0)

	p1 := maxProfit([]int{7, 1, 5, 7, 6, 4})
	fmt.Printf("p1 = %d \n\n", p1)

	p2 := maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Printf("p2 = %d \n\n", p2)

	p3 := maxProfit([]int{2, 4, 1})
	fmt.Printf("p3 = %d \n\n", p3)

	p4 := maxProfit([]int{2, 4})
	fmt.Printf("p4 = %d \n\n", p4)

	p5 := maxProfit([]int{2})
	fmt.Printf("p5 = %d \n\n", p5)

	p6 := maxProfit([]int{2, 5, 7, 8, 10, 8, 5, 7, 4, 2, 1})
	fmt.Printf("p6 = %d \n\n", p6)

	p7 := maxProfit([]int{7, 5, 7, 5, 7, 5, 5, 7})
	fmt.Printf("p7 = %d \n\n", p7)

	p8 := maxProfit([]int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9})
	fmt.Printf("p8 = %d \n\n", p8)
}
