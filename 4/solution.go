package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	allMid := (m + n + 1) / 2
	left := 0
	right := m
	maxLeft := 0
	minRight := 0
	num1Left := 0
	num1Right := 0
	num2Left := 0
	num2Right := 0

	for left <= right {
		i := (left + right) / 2
		j := allMid - i
		num1Left = math.MinInt
		if i != 0 {
			num1Left = nums1[i-1]
		}
		num1Right = math.MaxInt
		if i != m {
			num1Right = nums1[i]
		}
		num2Left = math.MinInt
		if j != 0 {
			num2Left = nums2[j-1]
		}
		num2Right = math.MaxInt
		if j != n {
			num2Right = nums2[j]
		}

		if maxLeft <= minRight {
			maxLeft = max(num1Left, num2Left)
			minRight = min(num1Right, num2Right)
			left = i + 1
		} else {
			right = i - 1
		}

	}

	if (m+n)%2 != 0 {
		return float64(maxLeft)
	} else {
		return (float64(maxLeft) + float64(minRight)) / 2
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
