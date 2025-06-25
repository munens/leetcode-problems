// https://leetcode.com/problems/spiral-matrix/description/
package main

import "fmt"

func reverse(arr []int) {
	j := len(arr) - 1
	midPoint := (len(arr) / 2) - 1
	for i := 0; i < len(arr); i++ {

		if i == j || i > midPoint {
			return
		}

		vi := arr[i]
		vj := arr[j]
		arr[i] = vj
		arr[j] = vi
		j--
	}
}

func canFinish(matrix [][]int, fin []int) bool {
	return len(fin) == (len(matrix) * len(matrix[0]))
}

func spiralOrderFunc(fin []int, matrix [][]int, dMin, dMax, xMin, xMax int) []int {
	fmt.Printf("dMin = %d, dMax = %d, xMin = %d, xMax = %d \n\n", dMin, dMax, xMin, xMax)

	// top row:
	topRow := matrix[dMin][xMin:xMax]
	fmt.Printf("topRow = %v \n", topRow)
	fin = append(fin, topRow...)

	if canFinish(matrix, fin) {
		return fin
	}

	// right column:
	var rightColumn []int
	for i := dMin + 1; i < dMax-1; i++ {
		rightColumn = append(rightColumn, matrix[i][xMax-1])
	}
	fmt.Printf("rightColumn = %v  \n", rightColumn)
	fin = append(fin, rightColumn...)

	// bottom row:
	bottomRow := matrix[dMax-1][xMin:xMax]
	reverse(bottomRow)
	fmt.Printf("bottomRow = %v  \n", bottomRow)
	fin = append(fin, bottomRow...)

	if canFinish(matrix, fin) {
		return fin
	}

	// left column:
	var leftColumn []int
	for i := dMax - 2; i > dMin; i-- {
		leftColumn = append(leftColumn, matrix[i][xMin])
	}
	fin = append(fin, leftColumn...)
	fmt.Printf("leftColumn = %v  \n", leftColumn)

	fmt.Printf("fin = %v  \n\n", fin)

	if canFinish(matrix, fin) {
		return fin
	} else {
		return spiralOrderFunc(fin, matrix, dMin+1, dMax-1, xMin+1, xMax-1)
	}
}

func spiralOrder(matrix [][]int) []int {
	return spiralOrderFunc([]int{}, matrix, 0, len(matrix), 0, len(matrix[0]))
}

func main() {

	matrix := [][]int{
		{1},
	}

	v := spiralOrder(matrix)
	fmt.Printf("	v = %v \n\n", v)

	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	v1 := spiralOrder(matrix1)
	fmt.Printf("v1 = %v \n\n", v1)

	matrix2 := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	}

	v2 := spiralOrder(matrix2)
	fmt.Printf("v2 = %v \n\n", v2)

	matrix3 := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
	}

	v3 := spiralOrder(matrix3)
	fmt.Printf("v3 = %v \n\n", v3)

	matrix4 := [][]int{
		{1, 2, 3, 4},
		{7, 8, 9, 10},
		{13, 14, 15, 16},
		{19, 20, 21, 22},
		{25, 26, 27, 28},
		{31, 32, 33, 34},
	}

	v4 := spiralOrder(matrix4)
	fmt.Printf("v4 = %v \n\n", v4)

	matrix5 := [][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}

	v5 := spiralOrder(matrix5)
	fmt.Printf("v5 = %v \n\n", v5)

}
