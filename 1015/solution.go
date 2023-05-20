package main

import "fmt"

func main() {
	fmt.Println(smallestRepunitDivByK(23))
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	dict := make(map[int]bool)
	x := 1
	i := 1
	for x%k != 0 {
		fmt.Println(x)
		if dict[x%k] {
			return -1
		}
		dict[x%k] = true
		i++
		x = x%k*10 + 1
	}
	return i
}
