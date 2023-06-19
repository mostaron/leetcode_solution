package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	fmt.Println(minNumberOfSemesters(13,
		[][]int{{12, 8}, {2, 4}, {3, 7}, {6, 8}, {11, 8}, {9, 4}, {9, 7}, {12, 4}, {11, 4}, {6, 4}, {1, 4}, {10, 7}, {10, 4}, {1, 7}, {1, 8}, {2, 7}, {8, 4}, {10, 8}, {12, 7}, {5, 4}, {3, 4}, {11, 7}, {7, 4}, {13, 4}, {9, 8}, {13, 8}},
		9))
}

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	need := make([]int, 1<<n)
	for _, edge := range relations {
		need[1<<(edge[1]-1)] |= 1 << (edge[0] - 1)
	}
	dp[0] = 0
	for i := 1; i < (1 << n); i++ {
		need[i] = need[i&(i-1)] | need[i&-i]
		if (need[i] | i) != i { // i 中有任意一门课程的前置课程没有完成学习
			continue
		}
		valid := i ^ need[i]                  // 当前学期可以进行学习的课程集合
		if bits.OnesCount(uint(valid)) <= k { // 如果个数小于 k，则可以全部学习，不再枚.举子集
			dp[i] = min(dp[i], dp[i^valid]+1)
		} else {
			for sub := valid; sub > 0; sub = (sub - 1) & valid {
				if bits.OnesCount(uint(sub)) <= k {
					dp[i] = min(dp[i], dp[i^sub]+1)
				}
			}
		}
	}
	return dp[(1<<n)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
