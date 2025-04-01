package main

import "fmt"

func merge(nums1 []int, nums2 []int, i int, j int, m int, n int) []int {

	if i == len(nums1)-1 || j == len(nums2) {
		return nums1
	}

	fmt.Printf("nums1 = %v, nums2 = %v, v = %d, i = %d, j = %d, m = %d \n\n", nums1, nums2, nums2[j], i, j, m)

	// handle case where nums1[i] <= nums2[j] < nums1[i+1]
	if nums1[i] <= nums2[j] && nums2[j] < nums1[i+1] {
		pre := nums1[0 : i+1]
		post := nums1[i+1 : len(nums1)-1]

		preCopy := append([]int(nil), pre...)
		preFin := append(preCopy, nums2[j])
		fmt.Printf("1st cond: pre = %v, preFin = %v, v = %d, post = %v \n\n", pre, preFin, nums2[j], post)
		fin := append(preFin, post...)

		return merge(fin, nums2, i+1, j+1, m+1, n)
	}

	// handle case where nums2[j] < nums1[i]
	if nums2[j] < nums1[i] {
		post := nums1[i : len(nums1)-1]
		fin := append([]int{nums2[j]}, post...)

		return merge(fin, nums2, i, j+1, m+1, n)
	}

	// handle case where nums2[j] > nums1[i]
	if nums2[j] >= nums1[i] && i == m-1 {
		pre := nums1[0 : i+1]
		post := nums1[i+2 : len(nums1)]

		preCopy := append([]int(nil), pre...)
		preFin := append(preCopy, nums2[j])
		fmt.Printf("3rd cond: pre = %v, preFin = %v, v = %d, post = %v \n\n", pre, preFin, nums2[j], post)
		fin := append(preFin, post...)

		return merge(fin, nums2, i+1, j+1, m+1, n)
	}

	return merge(nums1, nums2, i+1, j, m, n)
}

func mergeSortAlternative(nums1 []int, m int, nums2 []int, n int) []int {
	return merge(nums1, nums2, 0, 0, m, n)
}

func main() {

	//v0 := []int{}
	//mergeSort6(v0, 0, []int{1}, 1)
	//fmt.Println(v0)
	//fmt.Println("\n")
	//
	//v1 := []int{1, 2, 0}
	//mergeSort6(v1, 2, []int{}, 0)
	//fmt.Println(v1)
	//fmt.Println("\n")

	v2 := []int{1, 2, 3, 0, 0, 0}
	fmt.Printf("initial array, v2 = %v \n\n", v2)
	v2Fin := mergeSortAlternative(v2, 3, []int{2, 5, 6}, 3)
	fmt.Println(v2Fin)
	fmt.Println("\n")

	v3 := []int{3, 5, 7, 0, 0, 0, 0, 0}
	fmt.Printf("initial array, v3 = %v \n\n", v3)
	v3Fin := mergeSortAlternative(v3, 3, []int{2, 2, 4, 5, 9}, 5)
	fmt.Printf("final array, v3Fin = %v \n\n\n", v3Fin)

	v4 := []int{10, 11, 12, 15, 0, 0, 0, 0, 0, 0}
	fmt.Printf("initial array, v4 = %v \n\n", v4)
	v4Fin := mergeSortAlternative(v4, 4, []int{1, 2, 2, 8, 10, 15}, 6)
	fmt.Printf("final array, v3Fin = %v \n\n\n", v4Fin)

	v5 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	fmt.Printf("initial array, v5 = %v \n\n", v5)
	v5Fin := mergeSortAlternative(v5, 6, []int{1, 2, 2}, 3)
	fmt.Printf("final array, v3Fin = %v \n\n\n", v5Fin)

	v6 := []int{1, 0}
	fmt.Printf("initial array, v6 = %v \n\n", v6)
	v6Fin := mergeSortAlternative(v6, 1, []int{2}, 1)
	fmt.Printf("final array, v3Fin = %v \n\n\n", v6Fin)

	v7 := []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	fmt.Printf("initial array, v7 = %v \n\n", v7)
	v7Fin := mergeSortAlternative(v7, 5, []int{-1, -1, 0, 0, 1, 2}, 6)
	fmt.Printf("final array, v3Fin = %v \n\n\n", v7Fin)

}
