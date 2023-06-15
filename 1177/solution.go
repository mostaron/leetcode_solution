package main

func main() {

}

func canMakePaliQueries(s string, queries [][]int) []bool {
	result := make([]bool, len(queries))
	status := make([]int, len(s)+1)
	for i, c := range s {
		status[i+1] = status[i] ^ (1 << (c - 'a'))
	}
	for i, query := range queries {
		odd := 0
		current := status[query[1]+1] ^ status[query[0]]
		for current > 0 {
			if current%2 == 1 {
				odd++
			}
			current >>= 1
		}
		result[i] = odd <= query[2]*2+1
	}
	return result
}
