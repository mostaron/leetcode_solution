package main

import (
	"fmt"
	"math"
)

func main() {
	head := createLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	head = deleteDuplicates(head)
	printList(head)
	head = createLinkedList([]int{1, 2, 2})
	head = deleteDuplicates(head)
	printList(head)
	head = createLinkedList([]int{1, 1, 2})
	head = deleteDuplicates(head)
	printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println()
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head, p *ListNode
	for _, num := range nums {
		if p == nil {
			p = &ListNode{Val: num, Next: nil}
			head = p
		} else {
			p.Next = &ListNode{Val: num, Next: nil}
			p = p.Next
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHeader = &ListNode{Val: math.MinInt, Next: head}
	pointer := newHeader

	for pointer.Next != nil && pointer.Next.Next != nil {
		if pointer.Next.Val == pointer.Next.Next.Val {
			val := pointer.Next.Val
			for pointer.Next != nil && pointer.Next.Val == val {
				pointer.Next = pointer.Next.Next
			}
		} else {
			pointer = pointer.Next
		}
	}

	return newHeader.Next
}
