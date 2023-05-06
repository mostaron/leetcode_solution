package main

import "fmt"

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
	fmt.Println(minNumberOfFrogs("crcoakroak"))
	fmt.Println(minNumberOfFrogs("croakcrook"))
	fmt.Println(minNumberOfFrogs("croakcroa"))
	fmt.Println(minNumberOfFrogs("ccrrooaakk"))
}

func minNumberOfFrogs(croakOfFrogs string) int {
	frogs := make(map[rune]int)
	for _, c := range []rune(croakOfFrogs) {
		canSet := false
		switch {
		case c == 'c':
			if frogs['z'] == 0 {
				frogs[c]++
			} else {
				frogs[c]++
				frogs['z']--
			}

			canSet = true
		case c == 'r' && frogs['c'] > 0:
			frogs['c']--
			frogs['r']++
			canSet = true
		case c == 'o' && frogs['r'] > 0:
			frogs['r']--
			frogs['o']++
			canSet = true
		case c == 'a' && frogs['o'] > 0:
			frogs['o']--
			frogs['a']++
			canSet = true
		case c == 'k' && frogs['a'] > 0:
			frogs['a']--
			frogs['z']++
			canSet = true
		}
		if !canSet {
			return -1
		}
	}
	if frogs['c'] > 0 || frogs['r'] > 0 || frogs['o'] > 0 || frogs['a'] > 0 {
		return -1
	}
	return frogs['z']
}
