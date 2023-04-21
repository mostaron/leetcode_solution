package main

import (
	"fmt"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}

type LinkedList struct {
	size int
	head *Node
}
type Node struct {
	h, k        int
	back, front *Node
}

func (l *LinkedList) insertBefore(position, node *Node) {
	//fmt.Printf("insert [%v, %v], before [%v, %v]\n", node.h, node.k, position.h, position.k)
	node.front = position.front
	node.back = position
	if node.front != nil {
		node.front.back = node
	} else {
		l.head = node
	}
	position.front = node
	l.size++
}

func (l *LinkedList) insert(h, k int) {
	node := &Node{h: h, k: k, back: nil, front: nil}
	if l.size == 0 {
		l.head = node
		l.size++
		return
	}
	greater := 0
	var last *Node
	for cur := l.head; cur != nil; cur = cur.back {
		if cur.h >= h {
			if greater == k {
				l.insertBefore(cur, node)
				return
			}
			greater++
		}
		last = cur
	}
	last.back = node
	node.front = last
	l.size++

}
func (l *LinkedList) iterate() [][]int {
	result := make([][]int, l.size)
	i := 0
	for p := l.head; p != nil; p = p.back {
		//result = append(result, []int{p.h, p.k})
		result[i] = []int{p.h, p.k}
		i++
	}
	return result
}
func reconstructQueue(people [][]int) [][]int {
	l := LinkedList{}
	for k := 0; k < len(people); k++ {
		for i := 0; i < len(people); i++ {
			if people[i][1] == k {
				l.insert(people[i][0], people[i][1])
			}
		}
	}
	return l.iterate()
}
