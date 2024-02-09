// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func mergeSort(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 {
		nums1 = nums2
		return
	}

	if len(nums2) == 0 {
		nums1 = nums1[:m]
		return
	}

	var fin []int
	var prim []int
	var sec []int
	if nums1[0] <= nums2[0] {
		prim = nums1
		sec = nums2
	} else {
		prim = nums2
		sec = nums1
	}

	secMap := make(map[int][]int)

	for i := 0; i < len(sec); i++ {
		v := sec[i]
		_, ok := secMap[v]
		if ok {
			secMap[v] = append(secMap[v], v)
		} else {
			secMap[v] = []int{v}
		}
	}

	i := 0

	for {

		var primV int
		if i < m {
			primV = prim[i]
			fin = append(fin, primV)
		} else {
			primV = i
		}

		if v, ok := secMap[primV]; ok {
			fin = append(fin, v...)
		}

		if len(fin) == (m + n) {
			break
		} else {
			i++
		}
	}

	fmt.Println(nums1, fin)

	nums1 = fin

	return
}

func mergeSort2(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = nums2
	}

	if n == 0 {
		nums1 = nums1[:m]
	}

	nums2Map := make(map[int][]int)

	for i := 0; i < len(nums2); i++ {
		v := nums2[i]
		_, ok := nums2Map[v]
		if ok {
			nums2Map[v] = append(nums2Map[v], v)
		} else {
			nums2Map[v] = []int{v}
		}
	}

	i := 0
	j := 0
	for {
		if j > n-1 {
			break
		}

		v1 := nums1[i]
		v2 := nums2[j]

		nums := nums2Map[v2]
		a := len(nums)
		c := len(nums1) - a

		if v1 > v2 {
			pre := nums1[0:i:i]
			post := nums1[i:c]
			temp := append(pre, nums...)
			nums1 = append(temp, post...)
			i += a
			j += a
			continue
		}

		next := nums1[i+1]
		if v2 >= v1 && v2 < next || next == 0 {
			pre := nums1[0 : i+1 : i+1]
			post := nums1[i+1 : c]
			temp := append(pre, nums...)
			nums1 = append(temp, post...)
			i += a
			j += a
			continue
		}

		i += 1

	}

}

func mergeSort3(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		nums1 = nums2
		return nums1
	}

	if n == 0 {
		nums1 = nums1[:m]
		return nums1
	}

	nums2Map := make(map[int][]int)

	for i := 0; i < len(nums2); i++ {
		v := nums2[i]
		_, ok := nums2Map[v]
		if ok {
			nums2Map[v] = append(nums2Map[v], v)
		} else {
			nums2Map[v] = []int{v}
		}
	}

	i := 0
	j := 0
	for {
		if j > n-1 {
			break
		}

		v1 := nums1[i]
		v2 := nums2[j]

		nums := nums2Map[v2]
		a := len(nums)
		c := len(nums1) - a

		if v1 > v2 {
			pre := nums1[0:i:i]
			post := nums1[i:c]
			temp := append(pre, nums...)
			nums1 = append(temp, post...)
			i += a
			j += a
			continue
		}

		next := nums1[i+1]
		if v2 >= v1 && v2 < next || next == 0 {
			pre := nums1[0 : i+1 : i+1]
			post := nums1[i+1 : c]
			temp := append(pre, nums...)
			nums1 = append(temp, post...)
			i += a
			j += a
			continue
		}

		i += 1

	}

	return nums1

}

func mergeSort4(nums1 []int, m int, nums2 []int, n int) []int {
	var fin []int
	if m == 0 {
		fin = nums2
		return fin
	}

	if n == 0 {
		fin = nums1[:m]
		return fin
	}

	nums2Map := make(map[int][]int)

	for i := 0; i < len(nums2); i++ {
		v := nums2[i]
		_, ok := nums2Map[v]
		if ok {
			nums2Map[v] = append(nums2Map[v], v)
		} else {
			nums2Map[v] = []int{v}
		}
	}

	i := 0
	j := 0
	for {
		if j > n-1 {
			break
		}

		v1 := nums1[i]
		v2 := nums2[j]

		nums := nums2Map[v2]
		a := len(nums)
		c := len(nums1) - a

		if v1 > v2 {
			pre := nums1[0:i:i]
			post := nums1[i:c]
			temp := append(pre, nums...)
			fin = append(temp, post...)
			i += a
			j += a
			continue
		}

		next := nums1[i+1]
		if v2 >= v1 && v2 < next || next == 0 {
			pre := nums1[0 : i+1 : i+1]
			post := nums1[i+1 : c]
			temp := append(pre, nums...)
			fin = append(temp, post...)
			i += a
			j += a
			continue
		}

		i += 1

	}

	return fin

}

func mergeSort5(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1[0:], nums2)
		return
	}

	if n == 0 {
		nums1 = nums1[:m]
		return
	}

	nums2Map := make(map[int][]int)

	for i := 0; i < len(nums2); i++ {
		v := nums2[i]
		_, ok := nums2Map[v]
		if ok {
			nums2Map[v] = append(nums2Map[v], v)
		} else {
			nums2Map[v] = []int{v}
		}
	}

	fin := nums1

	i := 0
	j := 0

	count := 0

	for {
		if j == n {
			break
		}

		v1 := fin[i]
		v2 := nums2[j]

		nums := nums2Map[v2]
		a := len(nums)
		c := len(fin) - a

		fmt.Println("fin: ", fin, "i: ", i, "j: ", j, "nextIndex: ", i+1, ", v1: ", v1, ", v2: ", v2)
		if v1 > v2 {
			pre := fin[0:i:i]
			post := fin[i:c]
			temp := append(pre, nums...)
			fin = append(temp, post...)
			i += a
			j += a
			count += a
			continue
		}

		nextIndex := i + 1
		next := fin[nextIndex]
		if (v2 >= v1 && v2 < next) || (i+1 >= m+count) {
			pre := fin[0:nextIndex:nextIndex]
			post := fin[nextIndex:c]
			temp := append(pre, nums...)
			fin = append(temp, post...)
			i += a
			j += a
			count += a
			continue
		}

		i += 1
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = fin[i]
	}
}

func main() {

	v0 := []int{}
	mergeSort5(v0, 0, []int{1}, 1)
	fmt.Println(v0)
	fmt.Println("\n")

	v1 := []int{1, 2, 0}
	mergeSort5(v1, 2, []int{}, 0)
	fmt.Println(v1)
	fmt.Println("\n")

	v2 := []int{1, 2, 3, 0, 0, 0}
	mergeSort5(v2, 3, []int{2, 5, 6}, 3)
	fmt.Println(v2)
	fmt.Println("\n")

	v3 := []int{3, 5, 7, 0, 0, 0, 0, 0}
	mergeSort5(v3, 3, []int{2, 2, 4, 5, 9}, 5)
	fmt.Println(v3)
	fmt.Println("\n")

	v4 := []int{10, 11, 12, 15, 0, 0, 0, 0, 0, 0}
	mergeSort5(v4, 4, []int{1, 2, 2, 8, 10, 15}, 6)
	fmt.Println(v4)
	fmt.Println("\n")

	v5 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	mergeSort5(v5, 6, []int{1, 2, 2}, 3)
	fmt.Println(v5)
	fmt.Println("\n")

	v6 := []int{1, 0}
	mergeSort5(v6, 1, []int{2}, 1)
	fmt.Println(v6)
	fmt.Println("\n")

	v7 := []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	mergeSort5(v7, 5, []int{-1, -1, 0, 0, 1, 2}, 6)
	fmt.Println(v7)
	fmt.Println("\n")
}

// 1/31 look at this link - https://goplay.tools/snippet/hPR1gDr_Xkk
// notice that the slice can be modified from within the function. Read this - https://stackoverflow.com/questions/40514555/golang-passing-arrays-to-the-function-and-modifying-it
// It appears that when slices are passed around - and even in a function, by doing the assignment on the glo playground link just above or on this stack overflow link, we end
// up modifying the underlying array. It seems that I am really having this issue in this assignment here. I should probably change the way I am assigning. Start slow and then
// build from there..

// 2/3
// I should take my solution and break it up and explain main parts in Miro
// Then take 1 of the best solutions on leetcode and understand how it works here
// and break down how it works in Miro too!
