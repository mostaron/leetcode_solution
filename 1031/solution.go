package main

func main() {
	maxSumTwoNoOverlap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 3)
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	firstSum := make([]int, len(nums))
	secondSum := make([]int, len(nums))

	for index, num := range nums {
		if index == 0 {
			firstSum[index] = num
			secondSum[index] = num
			continue
		}
		if index < firstLen {
			firstSum[index] = num + firstSum[index-1]
		} else {
			firstSum[index] = num + firstSum[index-1] - nums[index-firstLen]
		}
		if index < secondLen {
			secondSum[index] = num + secondSum[index-1]
		} else {
			secondSum[index] = num + secondSum[index-1] - nums[index-secondLen]
		}
	}
	maxNum := 0
	for i := firstLen - 1; i < len(nums); i++ {
		for j := secondLen - 1; j < len(nums); j++ {
			if i-firstLen+1 <= j && i >= j-secondLen+1 {
				continue
			}
			maxNum = max(maxNum, firstSum[i]+secondSum[j])
		}
	}

	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
