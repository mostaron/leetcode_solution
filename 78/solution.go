package main

func main() {

}

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		tmp := [][]int{}
		for _, r := range result {
			n := []int{}
			n = append(n, r...)
			n = append(n, num)
			tmp = append(tmp, n)
		}
		result = append(result, tmp...)
	}
	return result
}
