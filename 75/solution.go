package main

import "fmt"

func main() {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	fmt.Println(colors)
}

func sortColors(nums []int) {
	colors := make(map[int]int)
	for _, num := range nums {
		colors[num]++
	}
	for i := 0; i < len(nums); i++ {
		if i < colors[0] {
			nums[i] = 0
		} else if i < colors[0]+colors[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
