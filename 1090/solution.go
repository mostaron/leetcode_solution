package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1))
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2))
	fmt.Println(largestValsFromLabels([]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 1))
	fmt.Println(largestValsFromLabels([]int{3, 2, 3, 2, 1}, []int{1, 0, 2, 2, 1}, 2, 1))
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	dict := make(map[int][]int)
	resultDict := make(map[int][]int)
	sum, added := 0, 0
	for index, value := range values {
		dict[value] = append(dict[value], labels[index])
	}
	//fmt.Println(dict)
	keys := []int{}
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	i := 1
	for added < numWanted && i <= len(keys) {
		key := keys[len(keys)-i]
		labs := dict[key]
		for j := 0; j < len(labs); j++ {
			//fmt.Println(labs[j], len(resultDict[labs[j]]))
			if len(resultDict[labs[j]]) >= useLimit {
				continue
			}

			sum += key
			added++
			resultDict[labs[j]] = append(resultDict[labs[j]], key)
			//fmt.Println(resultDict, len(resultDict), key)
			if added == numWanted {
				break
			}
		}
		i++
	}
	return sum
}
