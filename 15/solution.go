package main

import (
	"fmt"
)

func main() {
	test := []int{-1, 0, 1, 0}
	fmt.Println(threeSum(test))
}

func threeSum(nums []int) [][]int {
	tmp := make(map[int]map[int]int)
	maps := make(map[int][]int)
	result := [][]int{}
	for index, number := range nums {
		maps[number] = append(maps[number], index)
	}

	for k1, _ := range maps {
		for k2, _ := range maps {
			test := 0 - k1 - k2
			if len(maps[test]) > 0 {
				if ((k1 == k2 && k2 == test) && len(maps[k1]) >= 3) ||
					((k1 == k2 || k1 == test) && k2 != test && len(maps[k1]) >= 2) ||
					((k1 == k2 || k2 == test) && k1 != test && len(maps[k2]) >= 2) ||
					(k1 != k2 && k2 != test && k1 != test) {
					if k1 > k2 {
						k1, k2 = k2, k1
					}
					if k1 > test {
						k1, test = test, k1
					}
					if k2 > test {
						k2, test = test, k2
					}
					if tmp[k1] == nil {
						tmp[k1] = make(map[int]int)
					}
					tmp[k1][k2] = test
				}
			}
		}
	}

	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		test := 0 - nums[i] - nums[j]
	//		if len(maps[test]) > 0 {
	//			for _, k := range maps[test] {
	//				if k > i && k > j {
	//					arr := []int{nums[i], nums[j], nums[k]}
	//					sort.Ints(arr)
	//					if tmp[arr[0]] == nil {
	//						tmp[arr[0]] = make(map[int]int)
	//					}
	//					tmp[arr[0]][arr[1]] = arr[2]
	//				}
	//			}
	//		}
	//	}
	//}
	for k, v := range tmp {
		for k2, v2 := range v {
			result = append(result, []int{k, k2, v2})
		}
	}
	return result

}
