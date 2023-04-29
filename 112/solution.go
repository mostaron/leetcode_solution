package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	delta := targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		if delta == 0 {
			return true
		} else {
			return false
		}
	}

	result := false
	if root.Left != nil {
		result = result || hasPathSum(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		result = result || hasPathSum(root.Right, targetSum-root.Val)
	}
	return result
}
