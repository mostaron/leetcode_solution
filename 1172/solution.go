package main

import "sort"

func main() {

}

type StackNode struct {
	Val int
	Pre *StackNode
}

type Stack struct {
	Top  *StackNode
	Size int
}

type LinkedListNode struct {
	Val  *Stack
	Pre  *LinkedListNode
	Next *LinkedListNode
}

type DinnerPlates struct {
	capacity   int
	Head, Tail *LinkedListNode
	UnFull     map[int]*LinkedListNode
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		capacity: capacity,
		Head:     nil,
		Tail:     nil,
		UnFull:   make(map[int]*LinkedListNode),
	}
}

func (this *DinnerPlates) Push(val int) {
	if this.Head == nil {
		this.Head = &LinkedListNode{
			Val: &Stack{
				Top: &StackNode{
					Val: val,
				},
				Size: 1,
			},
			Pre:  nil,
			Next: nil,
		}
		this.Tail = this.Head
	} else {
		// find UnFull stack
		indexes := []int{}
		for k, v := range this.UnFull {
			if (v.Next != nil || v.Pre != nil) &&
				v.Val.Size < this.capacity {
				indexes = append(indexes, k)
			}
		}
		sort.Ints(indexes)

		stack := this.Tail
		// if exist then add node at first stack
		// else add at tail
		if len(indexes) > 0 {
			stack = this.UnFull[indexes[0]]
		}

		if stack.Val.Size == this.capacity {
			this.Tail = &LinkedListNode{
				Val: &Stack{
					Top: &StackNode{
						Val: val,
					},
					Size: 1,
				},
				Pre:  this.Tail,
				Next: nil,
			}
			this.Tail.Pre.Next = this.Tail
		} else {
			stack.Val.Top = &StackNode{
				Val: val,
				Pre: stack.Val.Top,
			}
			stack.Val.Size++

			if len(indexes) > 0 && stack.Val.Size == this.capacity {
				delete(this.UnFull, indexes[0])
			}
		}
	}

	//	var stack *LinkedListNode
	//	for stack = this.Head; stack != nil; stack = stack.Next {
	//		if stack.Val.Size < this.capacity {
	//			break
	//		}
	//	}
	//	if stack != nil {
	//		stack.Val.Top = &StackNode{
	//			Val: val,
	//			Pre: stack.Val.Top,
	//		}
	//		stack.Val.Size++
	//	} else {
	//		this.Tail = &LinkedListNode{
	//			Val: &Stack{
	//				Top: &StackNode{
	//					Val: val,
	//				},
	//				Size: 1,
	//			},
	//			Pre:  this.Tail,
	//			Next: nil,
	//		}
	//		this.Tail.Pre.Next = this.Tail
	//	}
	//}
}

func (this *DinnerPlates) Pop() int {
	if this.Head == nil {
		return -1
	}
	var stack = this.Tail
	if stack != nil {
		// record value
		val := stack.Val.Top
		// remove stack node
		stack.Val.Top = stack.Val.Top.Pre
		stack.Val.Size--
		// remove stack if needs
		this.removeTail()

		return val.Val
	}
	return -1
}

func (this *DinnerPlates) PopAtStack(index int) int {
	stack := this.Head
	for i := 0; i < index && stack != nil; i++ {
		stack = stack.Next
	}
	if stack == nil || stack.Val.Top == nil {
		return -1
	}

	// record value
	val := stack.Val.Top
	// remove stack node
	stack.Val.Top = stack.Val.Top.Pre
	stack.Val.Size--
	// remove stack if needs
	if stack == this.Tail {
		this.removeTail()
	} else {
		//mark unFull
		this.UnFull[index] = stack
	}
	return val.Val
}

func (this *DinnerPlates) removeTail() {
	for this.Tail != nil && this.Tail.Val.Size == 0 {
		p := this.Tail
		this.Tail = this.Tail.Pre
		p.Pre = nil
		if this.Tail == nil {
			this.Head = nil
		} else {
			this.Tail.Next = nil
		}
	}
}
