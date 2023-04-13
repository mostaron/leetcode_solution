package main

func main() {

}

func twoSum(nums []int, target int) []int {
	tmp := make(map[int][]int)
	for index, num := range nums {
		tmp[num] = append(tmp[num], index)
	}

	for k, v := range tmp {
		rest := target - k
		if len(tmp[rest]) > 0 && rest != k {
			return []int{v[0], tmp[rest][0]}
		} else if rest == k && len(v) >= 2 {
			return v[:2]
		}
	}
	return []int{}
}
