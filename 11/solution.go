package main

func main() {

}

func maxArea(height []int) int {
	i, j, a := 0, len(height)-1, 0
	for i < j {
		tmp := (j - i) * min(height[i], height[j])
		if a < tmp {
			a = tmp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return a
}

func min(h1, h2 int) int {
	if h1 < h2 {
		return h1
	}
	return h2
}
