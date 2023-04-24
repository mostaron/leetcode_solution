package main

import "fmt"

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root != nil && root.Left != nil {
		connects(root.Left, root.Right)
	}

	return root
}

func connects(nodes ...*Node) {
	if nodes[0] == nil {
		return
	}
	subNodes := make([]*Node, len(nodes)*2)
	i := 0
	for i < len(nodes)-1 {
		nodes[i].Next = nodes[i+1]
		subNodes[2*i] = nodes[i].Left
		subNodes[2*i+1] = nodes[i].Right
		i++
	}
	subNodes[2*i] = nodes[i].Left
	subNodes[2*i+1] = nodes[i].Right
	fmt.Println(subNodes)
	connects(subNodes...)
}
