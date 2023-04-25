package main

func main() {

}

type stackNode struct {
	ch  rune
	pre *stackNode
}

func isValid(s string) bool {
	var stack *stackNode
	for _, ch := range []rune(s) {
		if stack == nil {
			stack = &stackNode{ch: ch}
			continue
		}
		switch ch {
		case '(':
			stack = &stackNode{ch: ch, pre: stack}
			break
		case ')':
			if stack.ch != '(' {
				return false
			}
			stack = stack.pre
			break
		case '[':
			stack = &stackNode{ch: ch, pre: stack}
			break
		case ']':
			if stack.ch != '[' {
				return false
			}
			stack = stack.pre
			break
		case '{':
			stack = &stackNode{ch: ch, pre: stack}
			break
		case '}':
			if stack.ch != '{' {
				return false
			}
			stack = stack.pre
			break
		}
	}
	return stack == nil

}
