package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(38))
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
