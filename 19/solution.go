package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dict := make(map[int]*ListNode)
	i := 0
	p := head
	for p != nil {
		dict[i] = p
		i++
		p = p.Next
	}
	position := i - n
	if dict[position] == head {
		head = head.Next
	} else {
		dict[position-1].Next = dict[position].Next
	}
	return head
}
