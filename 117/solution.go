package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	subNodes := []*Node{}
	if root.Left != nil {
		subNodes = append(subNodes, root.Left)
	}
	if root.Right != nil {
		subNodes = append(subNodes, root.Right)
	}
	connects(subNodes...)

	return root
}

func connects(nodes ...*Node) {
	subNodes := []*Node{}
	i := 0
	for i < len(nodes)-1 {
		nodes[i].Next = nodes[i+1]
		if nodes[i].Left != nil {
			subNodes = append(subNodes, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			subNodes = append(subNodes, nodes[i].Right)
		}
		i++
	}
	if nodes[i].Left != nil {
		subNodes = append(subNodes, nodes[i].Left)
	}
	if nodes[i].Right != nil {
		subNodes = append(subNodes, nodes[i].Right)
	}
	if len(subNodes) == 0 {
		return
	}
	connects(subNodes...)
}
