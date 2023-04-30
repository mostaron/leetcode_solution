package main

func main() {

}

func numMovesStones(a int, b int, c int) []int {
	a, b = compAndEx(a, b)
	a, c = compAndEx(a, c)
	b, c = compAndEx(b, c)
	minMove, MaxMove := compAndEx(b-1-a, c-1-b)
	if minMove == 0 && MaxMove == 0 {
		return []int{0, 0}
	}
	if minMove == 0 || minMove == 1 {
		return []int{1, minMove + MaxMove}
	}
	return []int{2, minMove + MaxMove}

}

func compAndEx(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
