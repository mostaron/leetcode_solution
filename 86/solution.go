package main

import (
	"fmt"
)

func main() {
	head := partition(initNode([]int{1, 4, 3, 2, 5, 2}), 3)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initNode(arr []int) *ListNode {
	var first, p *ListNode
	for _, num := range arr {
		if first == nil {
			first = &ListNode{Val: num}
			p = first
		} else {
			p.Next = &ListNode{Val: num}
			p = p.Next
		}
	}
	return first
}

func partition(head *ListNode, x int) *ListNode {
	var less, lp *ListNode
	var great, gp *ListNode

	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			if lp == nil {
				less = p
			} else {
				lp.Next = p
			}
			lp = p
		} else {
			if gp == nil {
				great = p
			} else {
				gp.Next = p
			}
			gp = p
		}
	}
	if lp != nil {
		lp.Next = nil
	}
	if gp != nil {
		gp.Next = nil
	}

	if less == nil {
		return great
	} else {
		lp.Next = great
		return less
	}
}
