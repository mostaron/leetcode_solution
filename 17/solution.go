package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("2"))
}

var dict = map[uint8][]uint{
	'2': []uint{'a', 'b', 'c'},
	'3': []uint{'d', 'e', 'f'},
	'4': []uint{'g', 'h', 'i'},
	'5': []uint{'j', 'k', 'l'},
	'6': []uint{'m', 'n', 'o'},
	'7': []uint{'p', 'q', 'r', 's'},
	'8': []uint{'t', 'u', 'v'},
	'9': []uint{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	strs := []string{}
	if len(digits) == 0 {
		return strs
	}
	if len(digits) == 1 {
		for _, c := range dict[digits[0]] {
			strs = append(strs, string(c))
		}
		return strs
	} else {
		tmp := letterCombinations(digits[1:])
		for _, c := range dict[digits[0]] {
			for _, s := range tmp {
				strs = append(strs, string(c)+s)
			}
		}
	}
	return strs
}
