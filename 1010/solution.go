package main

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))
}

func numPairsDivisibleBy60(time []int) int {
	hash := make([]uint16, 60)
	count := 0
	for _, t := range time {
		tmp := t % 60
		count += int(hash[(60-tmp)%60])
		hash[tmp]++
	}
	return count
}
