package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
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
	merge(v0, 0, []int{1}, 1)
	fmt.Println(v0)
	fmt.Println("\n")

	v1 := []int{1, 2, 0}
	merge(v1, 2, []int{}, 0)
	fmt.Println(v1)
	fmt.Println("\n")

	v2 := []int{1, 2, 3, 0, 0, 0}
	merge(v2, 3, []int{2, 5, 6}, 3)
	fmt.Println(v2)
	fmt.Println("\n")

	v3 := []int{3, 5, 7, 0, 0, 0, 0, 0}
	merge(v3, 3, []int{2, 2, 4, 5, 9}, 5)
	fmt.Println(v3)
	fmt.Println("\n")

	v4 := []int{10, 11, 12, 15, 0, 0, 0, 0, 0, 0}
	merge(v4, 4, []int{1, 2, 2, 8, 10, 15}, 6)
	fmt.Println(v4)
	fmt.Println("\n")

	v5 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	merge(v5, 6, []int{1, 2, 2}, 3)
	fmt.Println(v5)
	fmt.Println("\n")

	v6 := []int{1, 0}
	merge(v6, 1, []int{2}, 1)
	fmt.Println(v6)
	fmt.Println("\n")

	v7 := []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	merge(v7, 5, []int{-1, -1, 0, 0, 1, 2}, 6)
	fmt.Println(v7)
	fmt.Println("\n")
}
