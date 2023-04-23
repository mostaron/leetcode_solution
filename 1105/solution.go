package main

import "math"

func main() {

}

func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books))
	for i := 0; i < len(books); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i, _ := range books {
		currentHeight, currentWidth := 0, 0
		for j := i; j >= 0; j-- {
			currentWidth += books[j][0]
			if currentWidth > shelfWidth {
				break
			}
			currentHeight = max(currentHeight, books[j][1])
			dp[i+1] = min(dp[i+1], dp[j]+currentHeight)
		}
	}
	return dp[len(books)-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
