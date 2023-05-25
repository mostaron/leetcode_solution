package main

import "fmt"

func main() {
	fmt.Println(oddString([]string{"adc", "wzy", "abc"}))
	fmt.Println(oddString([]string{"aaa", "bob", "ccc", "ddd"}))
}

func oddString(words []string) string {
	n := len(words[0])
	for i := 1; i < n; i++ {
		var a, b uint8
		var a1, b1 int
		a, b = 26, 26
		a1, b1 = 0, 0
		for index, word := range words {

			r := word[i] - word[i-1]
			//fmt.Println(a, b, a1, b1, r)
			if a == 26 {
				a = r
				a1 = index
				continue
			} else if a == r {
				if b != 26 {
					return words[b1]
				}
				continue
			} else if b == r {
				return words[a1]
			} else if index == len(words)-1 {
				return words[index]
			} else {
				b = r
				b1 = index
			}

		}
	}
	return ""
}
