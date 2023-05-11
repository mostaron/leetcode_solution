package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(2))
}
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, j, x, dir := 0, 0, 1, 0

	for x <= n*n {
		if result[i][j] == 0 {
			result[i][j] = x
			x++
		}

		if i+directions[dir][0] < 0 || i+directions[dir][0] >= n ||
			j+directions[dir][1] < 0 || j+directions[dir][1] >= n ||
			result[i+directions[dir][0]][j+directions[dir][1]] != 0 {
			dir = (dir + 1) % 4
		}
		i += directions[dir][0]
		j += directions[dir][1]
	}

	return result
}
