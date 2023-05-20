package main

func main() {

}

func maxValueAfterReverse(nums []int) int {
	value, n := 0, len(nums)
	for i := 0; i < n-1; i++ {
		value += Abs(nums[i] - nums[i+1])
	}
	mx1 := 0
	for i := 1; i < n-1; i++ {
		mx1 = Max(mx1, Abs(nums[0]-nums[i+1])-Abs(nums[i]-nums[i+1]))
		mx1 = Max(mx1, Abs(nums[n-1]-nums[i-1])-Abs(nums[i]-nums[i-1]))
	}
	mx2, mn2 := -100000, 100000
	for i := 0; i < n-1; i++ {
		x, y := nums[i], nums[i+1]
		mx2 = Max(mx2, Min(x, y))
		mn2 = Min(mn2, Max(x, y))
	}
	return value + Max(mx1, 2*(mx2-mn2))
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
