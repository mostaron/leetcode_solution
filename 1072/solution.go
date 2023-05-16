package main

import "fmt"

func main() {
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 1}, {1, 0}}))
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	dict := make(map[string]int)
	max := 0

	for _, row := range matrix {
		mask := row[0]
		tmp := byte(0)
		hash := []byte{}
		for i := 0; i < len(row); i += 8 {
			for j := 0; j < 8 && i+j < len(row); j++ {
				tmp += byte(row[i+j]^mask) << j
			}
			hash = append(hash, tmp)
		}
		dict[string(hash)]++
		if dict[string(hash)] > max {
			max = dict[string(hash)]
		}
	}

	return max
}
