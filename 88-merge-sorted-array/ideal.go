package main

import "fmt"

func mergeIdeal(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("nums1 = ", nums1, ", m = ", m, ", nums2 = ", nums2, ", n = ", n)
	for n != 0 {
		fmt.Println("m = ", m, ", n = ", n, ", m + n - 1 = ", m+n-1, ", nums1 = ", nums1, ", m != 0 && nums1[m-1] > nums2[n-1]", m != 0 && nums1[m-1] > nums2[n-1])
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func main() {
	//v0 := []int{}
	//mergeIdeal(v0, 0, []int{1}, 1)
	//fmt.Printf("%d \n", v0)
	//fmt.Println("\n")

	v1 := []int{1, 2, 0}
	mergeIdeal(v1, 2, []int{}, 0)
	fmt.Printf("%d \n\n", v1)

	v2 := []int{1, 2, 3, 0, 0, 0}
	mergeIdeal(v2, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%d \n\n", v2)

	v3 := []int{3, 5, 7, 0, 0, 0, 0, 0}
	mergeIdeal(v3, 3, []int{2, 2, 4, 5, 9}, 5)
	fmt.Printf("%d \n\n", v3)

	v4 := []int{10, 11, 12, 15, 0, 0, 0, 0, 0, 0}
	mergeIdeal(v4, 4, []int{1, 2, 2, 8, 10, 15}, 6)
	fmt.Printf("%d \n\n", v4)

	v5 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	mergeIdeal(v5, 6, []int{1, 2, 2}, 3)
	fmt.Printf("%d \n\n", v5)

	v6 := []int{1, 0}
	mergeIdeal(v6, 1, []int{2}, 1)
	fmt.Printf("%d \n\n", v6)

	v7 := []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	mergeIdeal(v7, 5, []int{-1, -1, 0, 0, 1, 2}, 6)
	fmt.Printf("%d \n\n", v7)
}
