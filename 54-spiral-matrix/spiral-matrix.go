// https://leetcode.com/problems/spiral-matrix/description/
package main

func reverse(arr []int) {
	j := 0
	midPoint := (len(arr) / 2) - 1
	for i := len(arr) - 1; i >= 0; i-- {

		if i == j {
			return
		}

		if j > midPoint {
			return
		}

		vi := arr[i]
		vj := arr[j]
		arr[i] = vj
		arr[j] = vi
		j++
	}
}

func spiralOrderFunc(fin []int, matrix [][]int, dMin, dMax, xMin, xMax int) {
	matrixSize := len(matrix) * len(matrix[0])
	if len(fin) == matrixSize {
		return
	}

	// top row:
	topRow := matrix[dMin][xMin : xMax-1]

	// right column:
	rightColumn := matrix[dMin : dMax-1][xMax]

	// bottom row:
	bottomRow := matrix[dMax][xMin+1 : xMax]
	reverse(bottomRow)

	//left column:
	leftColumn := matrix[dMin+2 : dMax][0]
	reverse(leftColumn)
}

func spiralOrder(matrix [][]int) {

	spiralOrderFunc([]int{}, matrix, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func main() {

}
