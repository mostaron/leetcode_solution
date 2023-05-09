package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n%2 == 1 {
		return x * myPow(x*x, n/2)
	} else {
		return myPow(x*x, n/2)
	}
}
