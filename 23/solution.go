package main

import "math"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	var nh, cur *ListNode
	var finished = 0
	for finished < len(lists) {
		mi, mn := 0, math.MaxInt
		var p *ListNode
		for index, node := range lists {
			if node == nil {
				continue
			}
			if node.Val < mn {
				mn = node.Val
				mi = index
				p = node
			}
		}
		if p == nil {
			break
		}
		lists[mi] = lists[mi].Next
		if lists[mi] == nil {
			finished++
		}
		if nh == nil {
			nh = p
			cur = p
		} else {
			cur.Next = p
			cur = cur.Next
		}
	}
	return nh
}
