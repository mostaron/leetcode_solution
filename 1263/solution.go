package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		minPushBox([][]byte{
			{'#', '.', '.', '#', 'T', '#', '#', '#', '#'},
			{'#', '.', '.', '#', '.', '#', '.', '.', '#'},
			{'#', '.', '.', '#', '.', '#', 'B', '.', '#'},
			{'#', '.', '.', '.', '.', '.', '.', '.', '#'},
			{'#', '.', '.', '.', '.', '#', '.', 'S', '#'},
			{'#', '.', '.', '#', '.', '#', '#', '#', '#'}}))
}

type Node struct {
	bx, by, sx, sy, step int
}

func minPushBox(grid [][]byte) int {

	xLen := len(grid)
	yLen := len(grid[0])
	bPath := make([][]bool, xLen*yLen)
	for index, _ := range bPath {
		bPath[index] = make([]bool, xLen*yLen)
	}
	node := Node{step: 0}
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				node.sx = i
				node.sy = j
			} else if c == 'B' {
				node.bx = i
				node.by = j
			}
		}
	}

	check := func(i, j int) bool {
		return i >= 0 && i < xLen && j >= 0 && j < yLen && grid[i][j] != '#'
	}
	queue := []Node{node}

	bPath[node.sx*yLen+node.sy][node.bx*yLen+node.by] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if grid[n.bx][n.by] == 'T' {
			return n.step
		}
		for _, dir := range dirs {
			sx, sy := n.sx+dir[0], n.sy+dir[1]
			if !check(sx, sy) {
				continue
			}
			if sx == n.bx && sy == n.by {
				bx, by := n.bx+dir[0], n.by+dir[1]
				if !check(bx, by) || bPath[sx*yLen+sy][bx*yLen+by] {
					continue
				}
				bPath[sx*yLen+sy][bx*yLen+by] = true
				//fmt.Printf("XXX set %d,%d to true \n", sx*yLen+sy, bx*yLen+by)
				queue = append(queue, Node{sx: sx, sy: sy, bx: bx, by: by, step: n.step + 1})
			} else if !bPath[sx*yLen+sy][n.bx*yLen+n.by] {
				bPath[sx*yLen+sy][n.bx*yLen+n.by] = true
				//fmt.Printf("NNN set %d,%d to true\n", sx*yLen+sy, n.bx*yLen+n.by)
				queue = append([]Node{{sx: sx, sy: sy, bx: n.bx, by: n.by, step: n.step}}, queue...)
			}
		}
		//fmt.Printf("%+v\n", queue)
	}
	return -1
}
