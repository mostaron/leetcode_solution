package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1463847412))
}

func abs(x int) (int32, int32) {
	if x < 0 {
		return -int32(x), -1
	} else {
		return int32(x), 1
	}
}

func reverse(x int) int {
	absX, flag := abs(x)
	result := int32(0)

	for absX > 0 {
		i := absX % 10
		absX /= 10
		tmp := result
		result = result*10 + i
		if tmp != result/10 {
			return 0
		}
	}
	return int(result) * int(flag)
}
