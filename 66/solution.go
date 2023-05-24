package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	n := len(digits) - 1
	digits[n]++
	for digits[n]/10 > 0 && n > 0 {
		digits[n] %= 10
		digits[n-1]++
		n--
	}

	if digits[0] >= 10 {
		digits[0] %= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
