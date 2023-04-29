package main

import "fmt"

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0}, {1, 1, 0}, {1, 1, 0},
	}))
}
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if (grid)[0][0] != 0 || (grid)[n-1][n-1] != 0 {
		return -1
	}
	if n == 1 && grid[0][0] == 0 {
		return 1
	}
	min := -1
	path := []stepNode{{0, 0}}
	min = minValid(min, shortestPath(&grid, path, 0, 1, 1))
	min = minValid(min, shortestPath(&grid, path, 1, 1, 1))
	min = minValid(min, shortestPath(&grid, path, 1, 0, 1))
	return min
}

func shortestPath(grid *[][]int, paths []stepNode, x, y, step int) int {
	if x < 0 || y < 0 {
		return -1
	}
	if isContain(paths, x, y) {
		return -1
	}
	if x >= len(*grid) || y >= len(*grid) {
		return -1
	}
	if x == len(*grid)-1 && y == len(*grid)-1 {
		//fmt.Println("!!!!", paths, x, y, step+1)
		return step + 1
	}
	if (*grid)[x][y] == 1 {
		return -1
	}
	paths = append(paths, stepNode{x, y})
	//fmt.Println(paths)
	min := -1
	min = minValid(min, shortestPath(grid, paths, x+1, y, step+1))
	min = minValid(min, shortestPath(grid, paths, x+1, y+1, step+1))
	min = minValid(min, shortestPath(grid, paths, x, y+1, step+1))
	min = minValid(min, shortestPath(grid, paths, x-1, y+1, step+1))
	min = minValid(min, shortestPath(grid, paths, x-1, y, step+1))
	min = minValid(min, shortestPath(grid, paths, x-1, y-1, step+1))
	min = minValid(min, shortestPath(grid, paths, x, y-1, step+1))
	min = minValid(min, shortestPath(grid, paths, x+1, y-1, step+1))

	return min

}

type stepNode struct {
	x, y int
}

func isContain(steps []stepNode, x, y int) bool {
	for _, step := range steps {
		if x == step.x && y == step.y {
			return true
		}
	}
	return false
}

func minValid(a, b int) int {
	//fmt.Println("?????", a, b)
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
