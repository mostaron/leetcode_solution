package main

func main() {

}

func averageValue(nums []int) int {
	total, count := 0, 0
	for _, num := range nums {
		if num%3 == 0 && num%2 == 0 {
			total += num
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return total / count
}
