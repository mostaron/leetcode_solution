package main

func main() {

}

func unequalTriplets(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				for k := j + 1; k < len(nums); k++ {
					if nums[k] != nums[i] && nums[k] != nums[j] {
						result++
					}
				}
			}
		}
	}
	return result
}
