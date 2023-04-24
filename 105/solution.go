package main

import "fmt"

func main() {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	position := find(inorder, preorder[0])
	root := &TreeNode{Val: preorder[0]}
	if position == -1 {
		return root
	}
	root.Left = buildTree(preorder[1:position+1], inorder[0:position])
	root.Right = buildTree(preorder[position+1:], inorder[position+1:])
	return root
}

func find(arr []int, target int) int {
	for index, num := range arr {
		if num == target {
			return index
		}
	}
	return -1
}
