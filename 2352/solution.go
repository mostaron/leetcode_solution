package main

import "fmt"

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}

func equalPairs(grid [][]int) int {
	count, n := 0, len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			isEqual := true
			for k := 0; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					isEqual = false
					break
				}
			}
			if isEqual {
				count++
			}
		}

	}
	return count
}
