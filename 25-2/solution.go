package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	p := head
	p.Next = &ListNode{Val: 2}
	p = p.Next
	p.Next = &ListNode{Val: 3}
	p = p.Next
	p.Next = &ListNode{Val: 4}
	p = p.Next
	p.Next = &ListNode{Val: 5}
	newHead := reverseKGroup(head, 2)
	fmt.Println(newHead.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	cur := head
	var newHead *ListNode
	var lastTail *ListNode
	for cur != nil {
		oldGroupHead := cur
		p := cur
		var oldGroupEnd, newGroupHead, newGroupEnd *ListNode
		groupCount := 0
		for p != nil && groupCount < k {
			groupCount++
			node := &ListNode{Val: p.Val}
			if newGroupHead == nil {
				newGroupHead = node
				newGroupEnd = node
			} else {
				node.Next = newGroupHead
				newGroupHead = node
			}
			oldGroupEnd = p
			p = p.Next

		}

		if groupCount < k {
			if newHead == nil {
				newHead = oldGroupHead
			} else {
				lastTail.Next = oldGroupHead
			}
			break
		} else {
			if newHead == nil {
				newHead = newGroupHead
			} else {
				lastTail.Next = newGroupHead
			}
			lastTail = newGroupEnd
			cur = oldGroupEnd.Next
		}
	}
	return newHead
}
