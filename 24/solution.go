package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, cur *ListNode
	p := head
	for p != nil {
		p1 := p
		p2 := p.Next
		if p2 == nil {
			cur.Next = p1
			break
		}
		p = p2.Next
		if newHead == nil {
			newHead = p2
			newHead.Next = p1
			cur = p1
			cur.Next = nil
		} else {
			cur.Next = p2
			cur = cur.Next
			cur.Next = p1
			cur = p1
			cur.Next = nil
		}
	}
	return newHead
}
