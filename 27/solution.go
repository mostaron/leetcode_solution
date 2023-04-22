package main

func main() {

}

func removeElement(nums []int, val int) int {
	removed := 0
	for index, num := range nums {
		if num != val {
			nums[index-removed] = num
		} else {
			removed++
		}
	}
	return len(nums) - removed
}
