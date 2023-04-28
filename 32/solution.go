package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

type Node struct {
	Val int
	Pre *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) pop() {
	if s.Top != nil {
		s.Top = s.Top.Pre
	}
}
func (s *Stack) push(val int) {
	node := &Node{
		Val: val,
	}
	if s.Top == nil {
		s.Top = node
		s.Top.Pre = nil
	} else {
		node.Pre = s.Top
		s.Top = node
	}

}

func longestValidParentheses(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	mask := make([]bool, len(s))
	stack := Stack{
		Top: nil,
	}
	for index, c := range []rune(s) {
		if c == '(' {
			stack.push(index)
		} else {
			if stack.Top != nil {
				mask[stack.Top.Val] = true
				mask[index] = true
				stack.pop()
			}
		}
	}
	maxLen, curLen := 0, 0
	for _, v := range mask {
		if v {
			curLen++
			maxLen = max(curLen, maxLen)
		} else {
			curLen = 0
		}
	}
	return maxLen

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
