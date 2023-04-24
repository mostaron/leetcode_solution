package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	newHead := rotateRight(head, 1)
	fmt.Println(newHead.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	last := head
	count := 1
	for last.Next != nil {
		count++
		last = last.Next
	}
	if count == 1 {
		return head
	}
	shift := k % count
	last.Next = head
	p := head
	for i := 0; i < count-shift-1; i++ {
		p = p.Next
	}
	head = p.Next
	p.Next = nil
	return head
}
