package main

import "fmt"

func main() {
	n := 1
	paths := [][]int{}
	fmt.Println(gardenNoAdj(n, paths))
}

type Node struct {
	flower int
	link   []*Node
}

func gardenNoAdj(n int, paths [][]int) []int {
	// create map and graph model
	tops := make(map[int]*Node)
	var result []int
	for _, path := range paths {
		top1 := tops[path[0]]
		if top1 == nil {
			top1 = &Node{}
			tops[path[0]] = top1
		}
		top2 := tops[path[1]]
		if top2 == nil {
			top2 = &Node{}
			tops[path[1]] = top2
		}
		top1.link = append(top1.link, top2)
		top2.link = append(top2.link, top1)
	}

	// iterate each tops, assign flower
	for i := 1; i <= n; i++ {
		node := tops[i]
		//no link to this node, create it and assign 1
		if node == nil {
			node = &Node{flower: 1}
			tops[i] = node
			result = append(result, 1)
			continue
		}
		// gather all linked node's flowers
		flowers := [4]bool{false, false, false, false}
		for _, link := range node.link {
			if link.flower != 0 {
				flowers[link.flower-1] = true
			}
		}
		// find empty flower and assign to this node
		var emptyFlower int
		for j := 0; j < 4; j++ {
			if flowers[j] == false {
				emptyFlower = j + 1
				break
			}
		}
		(*node).flower = emptyFlower
		result = append(result, emptyFlower)
	}

	return result
}
