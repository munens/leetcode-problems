package main

import "fmt"

func sum(arr []int) int {
	acc := 0
	for i := 0; i < len(arr); i++ {
		acc += arr[i]
	}

	return acc
}

func canCompleteCircuitIdeal(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	startingIdx := 0
	acc := 0

	for i := 0; i < len(gas); i++ {
		acc += gas[i] - cost[i]
		if acc < 0 {
			acc = 0
			startingIdx = i + 1
		}
	}

	return startingIdx
}

func main() {
	g0 := []int{1, 2, 3, 4, 5}
	c0 := []int{3, 4, 5, 1, 2}
	v0 := canCompleteCircuitIdeal(g0, c0)
	fmt.Printf("gas = %v, cost = %v, v0 = %d \n\n", g0, c0, v0)

	g1 := []int{2, 3, 4}
	c1 := []int{3, 4, 3}
	v1 := canCompleteCircuitIdeal(g1, c1)
	fmt.Printf("gas = %v, cost = %v, v1 = %d \n\n", g1, c1, v1)
}
