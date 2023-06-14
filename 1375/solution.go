package main

func main() {

}

func numTimesAllBlue(flips []int) int {
	max, count, result := 0, 0, 0
	for _, flip := range flips {
		if flip > max {
			max = flip
		}
		count++
		if count == max {
			result++
		}
	}
	return result
}
