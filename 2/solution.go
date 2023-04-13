package main

import "fmt"

func main() {
	l1 := &ListNode{}
	l2 := &ListNode{}
	t1 := l1
	t2 := l2
	t1.Val = 9
	t1.Next = &ListNode{}
	t1 = t1.Next
	t1.Val = 9

	t2.Val = 1

	result := addTwoNumbers(l1, l2)

	for {
		if result == nil {
			break
		}
		fmt.Print(result, "->")
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	moreThanTen := false
	var header, current *ListNode

	for {

		if !(l1 != nil || l2 != nil || moreThanTen) {
			break
		}

		if header == nil {
			header = &ListNode{}
			current = header
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}
		i1 := 0
		i2 := 0
		if l1 != nil {
			i1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			i2 = l2.Val
			l2 = l2.Next
		}
		sum := i1 + i2
		if moreThanTen {
			sum += 1
		}
		if sum >= 10 {
			moreThanTen = true
			sum = sum % 10
		} else {
			moreThanTen = false
		}
		current.Val = sum

	}
	return header
}
