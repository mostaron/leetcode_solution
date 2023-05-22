package main

func main() {

}
func numTilePossibilities(tiles string) int {
	count := make(map[rune]int)
	for _, t := range tiles {
		count[t]++
	}
	return dfs(len(tiles), count) - 1
}

func dfs(i int, count map[rune]int) int {
	if i == 0 {
		return 1
	}
	res := 1
	for t := range count {
		if count[t] > 0 {
			count[t]--
			res += dfs(i-1, count)
			count[t]++
		}
	}
	return res
}
