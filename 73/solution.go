package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{1, 0, 3}}
	setZeroes(grid)
	fmt.Println(grid)
}

func setZeroes(matrix [][]int) {
	firstRow, firstColumn := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstColumn = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstColumn {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if firstRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

}
