package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'},
	}, "ABCEFSADEESE"))
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for a, row := range board {
		for b, cell := range row {
			letters := []byte(word)
			if cell == word[0] && canFind(board, [][]int{{a, b}}, make([]bool, m*n), letters) {
				return true
			}
		}
	}
	return false
}

func canFind(board [][]byte, positions [][]int, path []bool, word []byte) bool {
	fmt.Println(positions, word)
	if len(word) == 0 {
		return true
	}
	if len(positions) == 0 {
		return false
	}
	for _, p := range positions {
		a, b := p[0], p[1]
		if path[a*len(board[0])+b] {
			continue
		}
		if board[a][b] == word[0] {
			fmt.Println("found", word[0], "at", a, b)
			directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
			tmp := [][]int{}
			for _, d := range directions {
				if a+d[0] >= 0 && a+d[0] < len(board) && b+d[1] >= 0 && b+d[1] < len(board[0]) {
					tmp = append(tmp, []int{a + d[0], b + d[1]})
				}
			}
			path[a*len(board[0])+b] = true
			result := canFind(board, tmp, path, word[1:])
			if result {
				return true
			} else {
				path[a*len(board[0])+b] = false
			}

		}
	}
	return false
}
