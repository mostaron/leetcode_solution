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
		groupStart := cur
		groupEnd := cur
		groupSize := 1
		for groupSize < k && groupEnd != nil {
			groupEnd = groupEnd.Next
			groupSize++
		}
		if groupSize <= k && groupEnd == nil {
			// 已抵达末尾，不做翻转
			if newHead == nil {
				newHead = groupStart
			} else {
				lastTail.Next = groupStart
			}
			break
		}
		//开始翻转
		cur = groupEnd.Next

		p := groupStart
		groupTail := groupStart
		groupStart = groupStart.Next
		p.Next = nil
		for groupStart != groupEnd {
			tmp := groupStart.Next
			groupStart.Next = p
			p = groupStart
			groupStart = tmp
		}
		groupEnd.Next = p

		if newHead == nil {
			newHead = groupEnd
		} else {
			lastTail.Next = groupEnd
		}
		lastTail = groupTail
	}
	return newHead
}
