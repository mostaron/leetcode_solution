package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isvalid(root.Left, 0, root.Val, true, false) &&
		isvalid(root.Right, root.Val, 0, false, true)
}

func isvalid(root *TreeNode, min, max int, isInfiniteMin, isInfiniteMax bool) bool {
	if root == nil {
		return true
	}

	if isInfiniteMin && root.Val < max ||
		isInfiniteMax && root.Val > min ||
		isInfiniteMin == false && isInfiniteMax == false && root.Val > min && root.Val < max {
		return isvalid(root.Left, min, root.Val, isInfiniteMin, true) &&
			isvalid(root.Right, root.Val, max, false, isInfiniteMax)
	} else {
		return false
	}
}
