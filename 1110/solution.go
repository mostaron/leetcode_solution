package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	result := delNodes(root, []int{1, 3, 5})
	for _, r := range result {
		fmt.Println(r.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	return delNode(root, nil, to_delete)
}

func delNode(node, parent *TreeNode, toDelete []int) []*TreeNode {
	result := []*TreeNode{}
	if node == nil {
		return result
	}
	if isIn(node, toDelete) {
		if parent != nil {
			if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		result = append(result, delNode(node.Left, nil, toDelete)...)
		result = append(result, delNode(node.Right, nil, toDelete)...)
	} else {
		if parent == nil {
			result = append(result, node)
		}
		result = append(result, delNode(node.Left, node, toDelete)...)
		result = append(result, delNode(node.Right, node, toDelete)...)
	}

	return result
}

func isIn(node *TreeNode, toDelete []int) bool {
	for _, num := range toDelete {
		if node.Val == num {
			return true
		}
	}

	return false
}
