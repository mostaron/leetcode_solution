package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	dict := make(map[rune]int)
	dict['I'] = 1
	dict['V'] = 5
	dict['X'] = 10
	dict['L'] = 50
	dict['C'] = 100
	dict['D'] = 500
	dict['M'] = 1000
	lastNum := 0
	num := 0
	for _, r := range []rune(s) {
		num += dict[r]
		if lastNum != 0 && dict[r] > lastNum {
			num -= lastNum * 2
		}
		lastNum = dict[r]

	}
	return num
}
