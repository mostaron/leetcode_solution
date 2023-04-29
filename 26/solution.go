package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	removed, unique := 0, 0
	for index, num := range nums {
		if index == 0 {
			unique++
			continue
		}
		if num == nums[index-1-removed] {
			removed++
			continue
		}
		unique++
		nums[index-removed] = num
	}
	return unique
}
