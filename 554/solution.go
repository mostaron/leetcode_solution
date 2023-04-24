package main

func main() {

}
func leastBricks(wall [][]int) int {
	projector := make(map[int]int)
	maxBorderCount := 0
	for _, line := range wall {
		length := 0
		for i := 0; i < len(line)-1; i++ {
			length += line[i]
			projector[length]++
			if maxBorderCount < projector[length] {
				maxBorderCount = projector[length]
			}
		}
	}
	return len(wall) - maxBorderCount
}
