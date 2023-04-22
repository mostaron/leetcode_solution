package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	i := []int{1, 1, 2, 2, 1, 1}
	t := Constructor(i)
	fmt.Println(t.Query(0, 5, 4))
	fmt.Println(t.Query(0, 3, 3))
	fmt.Println(t.Query(2, 3, 2))
}

type MajorityChecker struct {
	arr     *[]int
	results map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	m := MajorityChecker{arr: &arr, results: make(map[int][]int)}
	m.init()
	return m
}
func (m *MajorityChecker) init() {
	for index, i := range *m.arr {
		if m.results[i] == nil {
			m.results[i] = []int{index}
		} else {
			m.results[i] = append(m.results[i], index)
		}
		//sort.Sort(sort.IntSlice(m.results[i]))
	}
	//fmt.Println(m.results)
}

func isCatchThreshold(arr *[]int, left int, right int, thershold int) bool {
	if len(*arr) < thershold {
		return false
	}
	count := 0
	lessThanLeft := sort.SearchInts(*arr, left)
	lessThanRight := sort.SearchInts(*arr, right)
	if lessThanRight < len(*arr) && (*arr)[lessThanRight] == right {
		count++
	}
	count = count + lessThanRight - lessThanLeft
	//for _, position := range *arr {
	//	if position >= left && position <= right {
	//		count++
	//		if count >= thershold {
	//			return true
	//		}
	//	}
	//}
	return count >= thershold
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for i := 0; i < 20; i++ {
		pointer := rand.Intn(right+1-left) + left
		value := (*this.arr)[pointer]
		positions := this.results[value]
		if isCatchThreshold(&positions, left, right, threshold) {
			return value
		}
	}
	//for k, v := range this.results {
	//	if isCatchThreshold(&v, left, right, threshold) {
	//		return k
	//	}
	//}
	return -1
}
