package main

import "fmt"

func main() {
	fmt.Println(hardestWorker(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}))
}

func hardestWorker(n int, logs [][]int) int {
	maxTime := logs[0][1]
	minId := logs[0][0]
	for i := 1; i < len(logs); i++ {
		if logs[i][1]-logs[i-1][1] > maxTime {
			maxTime = logs[i][1] - logs[i-1][1]
			minId = logs[i][0]
		} else if logs[i][1]-logs[i-1][1] == maxTime {
			minId = min(minId, logs[i][0])
		}
	}
	return minId
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
