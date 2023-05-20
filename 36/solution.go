package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	// check if each row and column is valid
	for i := 0; i < 9; i++ {
		mr := make(map[byte]bool)
		mc := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if mr[board[i][j]] {
					return false
				}
				mr[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if mc[board[j][i]] {
					return false
				}
				mc[board[j][i]] = true
			}
		}
	}
	// check if each sub board is valid
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			mb := make(map[byte]bool)
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if board[i+m][j+n] != '.' {
						if mb[board[i+m][j+n]] {
							return false
						}
						mb[board[i+m][j+n]] = true
					}
				}
			}
		}
	}
	return true
}
