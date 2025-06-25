package main

import (
	"fmt"
	"slices"
)

func hIndexIdeal(citations []int) int {
	fmt.Printf("citations: %v \n\n", citations)

	slices.Sort(citations)
	fmt.Printf("sorted: %v \n\n", citations)

	hIndex := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			hIndex++
		}
	}

	return hIndex
}

func main() {
	fmt.Println("--------")
	arr0 := []int{3, 0, 6, 1, 5}
	v0 := hIndexIdeal(arr0)

	fmt.Printf("hIndex of: %v == %d \n\n", arr0, v0)

	fmt.Println("--------")
	arr1 := []int{1, 3, 1}
	v1 := hIndexIdeal(arr1)

	fmt.Printf("hIndex of: %v == %d \n\n", arr1, v1)

	fmt.Println("--------")
	arr2 := []int{100}
	v2 := hIndexIdeal(arr2)

	fmt.Printf("hIndex of: %v == %d \n\n", arr2, v2)

	fmt.Println("--------")
	arr3 := []int{0}
	v3 := hIndexIdeal(arr3)

	fmt.Printf("hIndex of: %v == %d \n\n", arr3, v3)

	fmt.Println("--------")
	arr4 := []int{0, 0, 2}
	v4 := hIndexIdeal(arr4)

	fmt.Printf("hIndex of: %v == %d \n\n", arr4, v4)

	fmt.Println("--------")
	arr5 := []int{11, 15}
	v5 := hIndexIdeal(arr5)

	fmt.Printf("hIndex of: %v == %d \n\n", arr5, v5)

	fmt.Println("--------")
	arr6 := []int{0, 0, 4, 4}
	v6 := hIndexIdeal(arr6)

	fmt.Printf("hIndex of: %v == %d \n\n", arr6, v6)

	fmt.Println("--------")
	arr7 := []int{0, 1, 2}
	v7 := hIndexIdeal(arr7)

	fmt.Printf("hIndex of: %v == %d \n\n", arr7, v7)

	fmt.Println("--------")
	arr8 := []int{1, 4, 7, 9}
	v8 := hIndexIdeal(arr8)

	fmt.Printf("hIndex of: %v == %d \n\n", arr8, v8)
}
