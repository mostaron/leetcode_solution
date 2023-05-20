package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxTotalFruits([][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2))
	//fmt.Println(maxTotalFruits([][]int{
	//	{0, 15}, {1, 2}, {2, 7}, {3, 65}, {4, 14}, {5, 49}, {6, 5}, {7, 27}, {8, 71}, {9, 6}, {10, 62}, {11, 15}, {12, 24}, {13, 51}, {14, 22}, {15, 79},
	//	{17, 98}, {18, 46}, {19, 91}, {21, 42}, {22, 31}, {23, 29}, {24, 95}, {25, 96}, {27, 94}, {28, 28}, {30, 62}, {31, 63}, {32, 94}, {33, 27},
	//	{34, 60}, {35, 97}, {36, 1}, {37, 57}, {39, 8}, {40, 92}, {41, 86}, {42, 37}, {43, 48}, {44, 3}, {45, 70}, {46, 64}, {47, 9}, {49, 100},
	//	{50, 33}, {51, 70}, {52, 18}, {54, 37}, {55, 100}, {56, 61}, {57, 33}, {58, 10}, {59, 27}, {60, 37}, {61, 77}, {62, 59}, {64, 30}, {65, 7},
	//	{66, 57}, {67, 5}, {68, 57}, {69, 13}, {70, 15}, {71, 95}, {72, 19}, {73, 50}, {74, 33}, {75, 20}, {76, 72}, {77, 95}, {78, 9}, {79, 18},
	//	{80, 85}, {82, 95}, {84, 85}, {86, 14}, {87, 26}, {88, 68}, {91, 61}, {92, 24}, {93, 32}, {94, 29}, {95, 77}, {97, 100}, {99, 59}, {100, 67}},
	//	50, 125))
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	left, right, sum, result := 0, 0, 0, 0
	getStep := func(left, right, position int) int {
		if position > fruits[right][0] {
			return position - fruits[left][0]
		}
		if position < fruits[left][0] {
			return fruits[right][0] - position
		}
		return min(position-fruits[left][0], fruits[right][0]-position) + fruits[right][0] - fruits[left][0]
	}
	for right < len(fruits) {
		sum += fruits[right][1]
		for left <= right && getStep(left, right, startPos) > k {
			sum -= fruits[left][1]
			left++
		}
		result = max(result, sum)
		right++
	}

	return result
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
