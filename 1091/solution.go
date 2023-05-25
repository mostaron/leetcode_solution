package main

import "fmt"

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	dp := make([][]bool, m)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	next := [][]int{{0, 0}}
	dp[0][0] = true
	step := 1
	for len(next) != 0 {
		tmp := [][]int{}
		for _, cell := range next {
			i, j := cell[0], cell[1]
			if grid[i][j] == 1 {
				continue
			}
			if i == m-1 && j == n-1 {
				return step
			}
			for _, direction := range directions {
				x, y := i+direction[0], j+direction[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 && !dp[x][y] {
					tmp = append(tmp, []int{x, y})
					dp[x][y] = true
				}
			}
		}
		next = tmp
		step++
	}
	return -1
}
