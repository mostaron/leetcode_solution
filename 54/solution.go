package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(matrix), len(matrix[0])
	dir, i, j := 0, 0, 0
	border := 0
	result := []int{}
	for len(result) < m*n {
		result = append(result, matrix[i][j])
		if !(dir == 0 && j < n-border-1) &&
			!(dir == 1 && i < m-border-1) &&
			!(dir == 2 && j > border) &&
			!(dir == 3 && i > border+1) {
			dir++
			border += dir / 4
			dir %= 4

		}
		i += directions[dir][0]
		j += directions[dir][1]
	}
	return result
}
