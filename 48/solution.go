package main

func main() {

}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1] = matrix[j][n-i-1], matrix[i][j]
			matrix[i][j], matrix[n-i-1][n-j-1] = matrix[n-i-1][n-j-1], matrix[i][j]
			matrix[i][j], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j]
		}
	}
}
