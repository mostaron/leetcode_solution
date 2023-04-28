package main

func main() {

}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 {
		if nums[i] < nums[j] {
			break
		}
		i--
		j--
	}

	for j = len(nums) - 1; j > i && i >= 0; j-- {
		if nums[i] < nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
			x := i + 1
			y := len(nums) - 1
			for x < y {
				nums[x], nums[y] = nums[y], nums[x]
				x++
				y--
			}

			return
		}
	}

	i, j = 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
