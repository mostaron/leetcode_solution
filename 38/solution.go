package main

import "fmt"

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	result := "1"
	for x := 1; x < n; x++ {
		count := rune(0)
		letter := '0'
		tmp := []rune{}
		for i, c := range []rune(result) {
			if i == 0 {
				count = 1
				letter = c
				continue
			}
			if letter == c {
				count++
			} else {
				tmp = append(tmp, count+'0')
				tmp = append(tmp, letter)
				count = 1
				letter = c
			}
		}
		tmp = append(tmp, count+'0')
		tmp = append(tmp, letter)
		result = string(tmp)
	}

	return result
}
