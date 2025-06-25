// https://leetcode.com/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func merge(left []int, right []int) []int {
	var merged []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// add remaining elements from left
	for i < len(left) {
		merged = append(merged, left[i])
		i++
	}

	// add remaining elements from right
	for j < len(right) {
		merged = append(merged, right[j])
		j++
	}

	return merged
}

func mergeSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	midPoint := len(arr) / 2

	left := mergeSort(arr[0:midPoint])
	right := mergeSort(arr[midPoint:])

	fmt.Printf("left: %v, right: %v \n\n", left, right)

	return merge(left, right)
}

func hIndex(citations []int) int {
	count := len(citations)

	// sort array using an algorithm
	sorted := mergeSort(citations)
	fmt.Printf("sorted: %v \n\n", sorted)

	if sorted[count-1] == 0 {
		return 0
	}

	if count == 1 {
		return 1
	}

	res := sorted[0]
	if res > count {
		res = count
	}

	// find h-index on sorted array
	// - traverse backwards in array
	// - find no. of elements above certain threshold
	for i := count - 1; i >= 0; i-- {
		vi := sorted[i]

		viewed := len(sorted[i:])
		fmt.Printf("res = %d, v = %d, viewed = %d \n\n", res, vi, viewed)
		if vi == viewed {
			res = vi
			break
		}

		if vi < viewed {
			res = len(sorted[i+1:])
			break
		}
	}

	return res
}

func main() {
	fmt.Println("--------")
	arr0 := []int{3, 0, 6, 1, 5}
	v0 := hIndex(arr0)

	fmt.Printf("hIndex of: %v == %d \n\n", arr0, v0)

	fmt.Println("--------")
	arr1 := []int{1, 3, 1}
	v1 := hIndex(arr1)

	fmt.Printf("hIndex of: %v == %d \n\n", arr1, v1)

	fmt.Println("--------")
	arr2 := []int{100}
	v2 := hIndex(arr2)

	fmt.Printf("hIndex of: %v == %d \n\n", arr2, v2)

	fmt.Println("--------")
	arr3 := []int{0}
	v3 := hIndex(arr3)

	fmt.Printf("hIndex of: %v == %d \n\n", arr3, v3)

	fmt.Println("--------")
	arr4 := []int{0, 0, 2}
	v4 := hIndex(arr4)

	fmt.Printf("hIndex of: %v == %d \n\n", arr4, v4)

	fmt.Println("--------")
	arr5 := []int{11, 15}
	v5 := hIndex(arr5)

	fmt.Printf("hIndex of: %v == %d \n\n", arr5, v5)

	fmt.Println("--------")
	arr6 := []int{0, 0, 4, 4}
	v6 := hIndex(arr6)

	fmt.Printf("hIndex of: %v == %d \n\n", arr6, v6)

	fmt.Println("--------")
	arr7 := []int{0, 1, 2}
	v7 := hIndex(arr7)

	fmt.Printf("hIndex of: %v == %d \n\n", arr7, v7)

	fmt.Println("--------")
	arr8 := []int{1, 4, 7, 9}
	v8 := hIndex(arr8)

	fmt.Printf("hIndex of: %v == %d \n\n", arr8, v8)
}
