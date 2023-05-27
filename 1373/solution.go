package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 1}
	root.Left.Left.Left = &TreeNode{Val: 9}
	root.Left.Right.Left = &TreeNode{Val: -5}
	root.Left.Right.Left.Right = &TreeNode{Val: -3}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Left.Right.Right.Right = &TreeNode{Val: 10}

	fmt.Println(maxSumBST(root))
	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 10}
	root.Right.Left = &TreeNode{Val: -5}
	root.Right.Right = &TreeNode{Val: 20}
	fmt.Println(maxSumBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int, int)
	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt, 0
		}

		lMin, lMax, lSum := dfs(node.Left)  // 递归左子树
		rMin, rMax, rSum := dfs(node.Right) // 递归右子树
		x := node.Val
		if x <= lMax || x >= rMin { // 不是二叉搜索树
			return math.MinInt, math.MaxInt, 0
		}

		s := lSum + rSum + x // 这棵子树的所有节点值之和
		ans = max(ans, s)

		return min(lMin, x), max(rMax, x), s
	}
	dfs(root)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
