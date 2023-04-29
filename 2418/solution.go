package main

import "sort"

func main() {

}
func sortPeople(names []string, heights []int) []string {
	dict := make(map[int]string)
	for index, height := range heights {
		dict[height] = names[index]
	}
	sort.Ints(heights)
	newNames := make([]string, len(names))
	for i := len(heights) - 1; i >= 0; i-- {
		newNames[len(names)-1-i] = dict[heights[i]]
	}
	return newNames
}
