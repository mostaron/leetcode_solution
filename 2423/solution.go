package main

import "math"

func main() {

}
func equalFrequency(word string) bool {
	countDict := make(map[rune]int)
	maxCount := 0
	for _, c := range []rune(word) {
		countDict[c]++
		maxCount = max(maxCount, countDict[c])
	}
	if maxCount == 1 {
		return true
	}
	if len(countDict) == 1 {
		return true
	}
	revMap := make(map[int][]rune)
	mi := math.MaxInt
	for k, v := range countDict {
		if len(revMap[v]) == 0 {
			revMap[v] = []rune{k}
		} else {
			revMap[v] = append(revMap[v], k)
		}
		mi = min(mi, v)
	}
	if len(revMap) == 1 && len(revMap[1]) != 0 {
		return true
	}
	if len(revMap) == 2 {
		if len(revMap[1]) == 1 {
			return true
		}
		if len(revMap[maxCount]) == 1 && maxCount-mi == 1 {
			return true
		}
	}
	return false

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
