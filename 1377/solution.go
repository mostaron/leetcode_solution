package main

import "fmt"

func main() {
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 1, 7))
	fmt.Println(frogPosition(3, [][]int{{2, 1}, {3, 2}}, 1, 2))
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 20, 6))
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	links := make([][]int, n+1)
	for _, l := range edges {
		links[l[0]] = append(links[l[0]], l[1])
		links[l[1]] = append(links[l[1]], l[0])
	}
	path := make(map[int]bool)
	path[1] = true
	return jump(links, 1, t, target, 1, path)
}

func jump(links [][]int, step, t, target int, pos float64, path map[int]bool) float64 {
	next := []int{}
	for _, l := range links[step] {
		if path[l] {
			continue
		}
		next = append(next, l)
	}
	if t == 0 || len(next) == 0 {
		if step != target {
			return float64(0)
		} else {
			return pos
		}
	}

	for _, l := range next {

		path[l] = true
		result := jump(links, l, t-1, target, pos*(1.0/float64(len(next))), path)

		if result != 0 {
			return result
		}
		delete(path, l)
	}
	return 0

}
