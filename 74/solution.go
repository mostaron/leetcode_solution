package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 2))
}

func searchMatrix(matrix [][]int, target int) bool {
	n, size := len(matrix[0]), len(matrix)*len(matrix[0])
	i, j := 0, size
	for i <= j {
		mid := (i + j) / 2
		if mid/n >= len(matrix) {
			return false
		}
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return false
}
