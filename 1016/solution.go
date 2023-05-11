package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(queryString("0110", 3))
	//fmt.Println(queryString("0110", 4))
	//fmt.Println(queryString("011010101010111101010101011111111111111111111111111111111110000000000000011111101010101"+
	//	"00101010101010101010101010111101010101011111111111111111111111111111111110000000000000011111101010101001010101"+
	//	"0101010101010100", 1000000000))
	fmt.Println(isContain("1111", "110"))
}

func queryString(s string, n int) bool {
	num, _ := strconv.ParseInt(s, 2, 0)
	if int(num) < n {
		return false
	}
	for i := 1; i <= n; i++ {
		strI := strconv.FormatInt(int64(i), 2)
		if isContain(s, strI) {
			continue
		}
		return false
	}
	return true
}
func isContain(s, query string) bool {
	i, j := 0, 0
	for i <= len(s)-len(query) && j < len(query) {
		if s[i+j] == query[j] {
			j++
		} else {
			i++
			j = 0
		}
	}
	return j >= len(query)

}
