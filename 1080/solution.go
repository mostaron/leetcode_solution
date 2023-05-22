package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	isSuficient := pathSufficient(root, limit, 0)
	if !isSuficient {
		return nil
	}
	return root
}

func pathSufficient(root *TreeNode, limit, sum int) bool {
	isRemove := false
	if root.Left != nil && !pathSufficient(root.Left, limit, sum+root.Val) {
		root.Left = nil
		isRemove = true
	}
	if root.Right != nil && !pathSufficient(root.Right, limit, sum+root.Val) {
		root.Right = nil
		isRemove = true
	}
	if isRemove && root.Left == nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum+root.Val >= limit
	}
	return true
}

//func sufficientSubset(root *TreeNode, limit int) *TreeNode {
//	isSuficient := pathSufficient(root, nil, limit, 0)
//	if !isSuficient {
//		return nil
//	}
//	return root
//}
//
//func pathSufficient(node, parent *TreeNode, limit, sum int) bool {
//
//	isRemoved := false
//	if node.Left != nil && !pathSufficient(node.Left, node, limit, sum+node.Val) {
//		node.Left = nil
//		isRemoved = true
//	}
//	if node.Right != nil && !pathSufficient(node.Right, node, limit, sum+node.Val) {
//		node.Right = nil
//		isRemoved = true
//	}
//
//	if isRemoved && parent != nil {
//		direction := 0
//		if parent.Right == node {
//			direction = 1
//		}
//
//		if node.Left == nil && node.Right == nil {
//			if direction == 0 {
//				parent.Left = nil
//			} else {
//				parent.Right = nil
//			}
//		} else if node.Left != nil {
//			if direction == 0 {
//				parent.Left = node.Left
//			} else {
//				parent.Right = node.Left
//			}
//		} else if node.Right != nil {
//			if direction == 0 {
//				parent.Left = node.Right
//			} else {
//				parent.Right = node.Right
//			}
//		}
//	}
//
//	if node.Left == nil && node.Right == nil {
//		return sum+node.Val > limit
//	}
//	return true
//}
