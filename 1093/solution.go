package main

import "fmt"

func main() {
	fmt.Println(sampleStats([]int{0, 4, 3, 2, 2}))
}

func sampleStats(count []int) []float64 {
	min, max, total, totalCount, modeNum, mode := -1, -1, 0, 0, 0, 0
	for num, c := range count {
		if c == 0 {
			continue
		}
		if min == -1 {
			min = num
		}
		max = num
		totalCount += c
		total += num * c
		if c > modeNum {
			modeNum = c
			mode = num
		}
	}
	var middle []int
	if totalCount%2 == 0 {
		middle = []int{totalCount / 2, totalCount/2 + 1}
	} else {
		middle = []int{totalCount/2 + 1}
	}
	middleValue := []int{-1, -1}
	passed := 0
	for num, c := range count {
		if c == 0 {
			continue
		}
		if c+passed >= middle[0] && middleValue[0] == -1 {
			middleValue[0] = num
			if len(middle) == 1 {
				break
			}
		}
		if len(middle) == 2 && c+passed >= middle[1] {
			middleValue[1] = num
			break
		}
		passed += c
	}
	mid := 0
	for _, i := range middleValue {
		if i == -1 {
			continue
		}
		mid += i
	}
	return []float64{
		float64(min), float64(max), float64(total) / float64(totalCount),
		float64(mid) / float64(len(middle)), float64(mode),
	}
}
