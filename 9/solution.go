package main

func main() {

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	i := x
	j := 0
	for x > 0 {
		j = j*10 + x%10
		x /= 10
	}
	return i == j

}
