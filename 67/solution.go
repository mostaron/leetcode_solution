package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	m, n := []rune(a), []rune(b)
	if len(m) < len(n) {
		m, n = n, m
	}

	addition := uint8(0)
	for i := 1; i <= len(n) || addition != 0; i++ {
		d := addition
		if i <= len(a) {
			d += a[len(a)-i] - '0'
		}
		if i <= len(b) {
			d += b[len(b)-i] - '0'
		}
		if d == 3 {
			addition = 1
			d = 1
		} else if d == 2 {
			addition = 1
			d = 0
		} else {
			addition = 0
		}
		if i <= len(m) {
			m[len(m)-i] = rune(d + '0')
		} else {
			m = append([]rune{rune(d + '0')}, m...)
		}
	}
	if addition == 1 {
		m = append([]rune{'1'}, m...)
	}
	return string(m)
}
