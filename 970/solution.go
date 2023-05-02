package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(powerfulIntegers(3, 5, 15))
}
func powerfulIntegers(x int, y int, bound int) []int {
	if x == 1 && y == 1 {
		if x+y <= bound {
			return []int{2}
		} else {
			return []int{}
		}
	}
	resultMap := make(map[int]bool)

	if x > y {
		x, y = y, x
	}
	if x == 1 {
		for tmpy := 1; tmpy < bound; tmpy *= y {
			if x+tmpy <= bound {
				resultMap[x+tmpy] = true
			}
		}
	} else {
		for tmpx := 1; tmpx < bound; tmpx *= x {
			for tmpy := 1; tmpy < bound; tmpy *= y {
				if tmpx+tmpy <= bound {
					resultMap[tmpx+tmpy] = true
				}
			}
		}
	}
	result := []int{}
	for k, _ := range resultMap {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}
