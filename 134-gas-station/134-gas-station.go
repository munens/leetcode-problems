// https://leetcode.com/problems/gas-station/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func canCompleteCircuitRecursively(gas []int, cost []int, stationIdx int) int {

	n := len(gas)

	if stationIdx == n {
		return -1
	}

	total := 0

	i := stationIdx

	for {
		total = (total + gas[i]) - cost[i]

		fmt.Printf("i = %d, current total = %d, gas = %d, cost = %d, stationIdx = %d  \n\n", i, total, gas, cost, stationIdx)

		if total < 0 {
			return canCompleteCircuitRecursively(gas, cost, stationIdx+1)
		}

		i++

		if i == stationIdx || (i == n && stationIdx == 0) {
			return stationIdx
		}

		if i == n {
			i = 0
		}
	}

	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)

	i := 0
	startingIdx := -1

	for i < n {
		if gas[i] > cost[i] {
			startingIdx = i
			break
		} else {
			i++
		}
	}

	fmt.Printf("startingIdx = %d \n\n", startingIdx)

	if startingIdx == -1 {
		return -1
	}

	return canCompleteCircuitRecursively(gas, cost, startingIdx)
}

func main() {
	g0 := []int{1, 2, 3, 4, 5}
	c0 := []int{3, 4, 5, 1, 2}
	v0 := canCompleteCircuit(g0, c0)
	fmt.Printf("gas = %v, cost = %v, v0 = %d \n\n", g0, c0, v0)

	g1 := []int{2, 3, 4}
	c1 := []int{3, 4, 3}
	v1 := canCompleteCircuit(g1, c1)
	fmt.Printf("gas = %v, cost = %v, v1 = %d \n\n", g1, c1, v1)
}
