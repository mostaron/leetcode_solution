package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
}
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	rev := false
	if dividend > 0 {
		rev = !rev
		dividend = -dividend
	}
	if divisor > 0 {
		rev = !rev
		divisor = -divisor
	}
	if dividend > divisor {
		return 0
	}

	ans, shift := 0, 0
	for dividend-divisor<<shift < divisor<<shift {
		shift++
	}
	ans = 1<<shift + divide(dividend-divisor<<shift, divisor)
	if rev {
		return -ans
	}
	return ans
}
