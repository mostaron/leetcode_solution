package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i, row := range obstacleGrid {
		for j, grid := range row {
			if grid == 1 {
				dp[i][j] = 0

			} else if i == 0 || j == 0 {
				if j > 0 {
					dp[i][j] = min(1, dp[i][j-1])
				} else if i > 0 {
					dp[i][j] = min(1, dp[i-1][j])
				} else {
					dp[i][j] = 1
				}

			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
