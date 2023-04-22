package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	maxLeft := maxDiff(root.Left, []int{root.Val})
	maxRight := maxDiff(root.Right, []int{root.Val})
	if maxLeft > maxRight {
		return maxLeft
	} else {
		return maxRight
	}
}

func maxDiff(root *TreeNode, parents []int) int {
	if root == nil {
		return 0
	}

	maxDiffVal := 0
	for p := range parents {
		if maxDiffVal < abs(root.Val-p) {
			maxDiffVal = abs(root.Val - p)
		}
	}

	parents = append(parents, root.Val)
	maxLeft := maxDiff(root.Left, parents)
	maxRight := maxDiff(root.Right, parents)
	if maxLeft > maxDiffVal {
		maxDiffVal = maxLeft
	}
	if maxRight > maxDiffVal {
		maxDiffVal = maxRight
	}
	fmt.Println(maxDiffVal, parents)
	return maxDiffVal
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
