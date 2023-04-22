package main

func main() {

}

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ {
		maxNum := 0
		for j := i - 1; j >= max(0, i-k); j-- {
			maxNum = max(maxNum, arr[j])
			dp[i] = max(dp[i], dp[j]+maxNum*(i-j))
		}
	}
	return dp[len(arr)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
