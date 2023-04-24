package main

import "math"

func main() {

}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := Stack{bottom: nil}
	i, j := 0, 0

	for i < len(popped) {
		if stack.getLast() == popped[i] {
			stack.pop()
			i++
			continue
		}
		if j < len(pushed) {
			stack.push(pushed[j])
			j++
		} else {
			return false
		}
	}
	return true
}

type Node struct {
	val int
	pre *Node
}
type Stack struct {
	bottom *Node
}

func (s *Stack) getLast() int {
	if s.bottom == nil {
		return math.MinInt
	}
	return s.bottom.val
}
func (s *Stack) push(val int) {
	node := &Node{
		pre: s.bottom,
		val: val,
	}
	s.bottom = node
}

func (s *Stack) pop() {
	node := s.bottom
	s.bottom = s.bottom.pre
	node.pre = nil
}
