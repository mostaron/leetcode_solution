package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newHead, cur *ListNode
	for list1 != nil && list2 != nil {
		var p *ListNode
		if list1.Val < list2.Val {
			p = list1
			list1 = list1.Next
		} else {
			p = list2
			list2 = list2.Next
		}
		p.Next = nil
		if newHead == nil {
			newHead = p
			cur = p
		} else {
			cur.Next = p
			cur = cur.Next
		}
	}
	if list1 != nil {
		if cur != nil {
			cur.Next = list1
		} else {
			newHead = list1
		}
	} else if list2 != nil {
		if cur != nil {
			cur.Next = list2
		} else {
			newHead = list2
		}
	}
	return newHead
}
