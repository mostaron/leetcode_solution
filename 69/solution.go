package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(15))
	fmt.Println(mySqrt(80))
}

func mySqrt(x int) int {
	return recSqrt(x, 0)
}

func recSqrt(x, result int) int {
	inc := 1
	for (inc+result)*(inc+result) <= x {
		inc <<= 1
	}
	inc >>= 1
	if inc == 0 {
		return result
	}
	result += inc
	return recSqrt(x, result)
}
