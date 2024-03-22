package main

import "fmt"

func maxProfitIdeal(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for day := 1; day < len(prices); day++ {
		if buyPrice > prices[day] {
			buyPrice = prices[day]
		}
		profit := prices[day] - buyPrice
		if profit > 0 {
			maximumProfit += profit
			buyPrice = prices[day]
		}
	}
	return maximumProfit
}

func main() {
	p0 := maxProfitIdeal([]int{7, 1, 5, 3, 6, 4})
	fmt.Printf("p0 = %d \n\n", p0)

	p1 := maxProfitIdeal([]int{7, 1, 5, 7, 6, 4})
	fmt.Printf("p1 = %d \n\n", p1)

	p2 := maxProfitIdeal([]int{7, 6, 4, 3, 1})
	fmt.Printf("p2 = %d \n\n", p2)

	p3 := maxProfitIdeal([]int{2, 4, 1})
	fmt.Printf("p3 = %d \n\n", p3)

	p4 := maxProfitIdeal([]int{2, 4})
	fmt.Printf("p4 = %d \n\n", p4)

	p5 := maxProfitIdeal([]int{2})
	fmt.Printf("p5 = %d \n\n", p5)

	p6 := maxProfitIdeal([]int{2, 5, 7, 8, 10, 8, 5, 7, 4, 2, 1})
	fmt.Printf("p6 = %d \n\n", p6)

	p7 := maxProfitIdeal([]int{7, 5, 7, 5, 7, 5, 5, 7})
	fmt.Printf("p7 = %d \n\n", p7)

	p8 := maxProfitIdeal([]int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9})
	fmt.Printf("p8 = %d \n\n", p8)
}
