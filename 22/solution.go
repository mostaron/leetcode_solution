package main

import (
	"fmt"
	"sort"
)

/*
*
NOT FIXED
*/
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	} else {
		last := generateParenthesis(n - 1)
		result := []string{}
		for _, str := range last {
			result = append(result, "("+str+")")
			if ("()" + str) != (str + "()") {
				result = append(result, "()"+str)
			}
			result = append(result, str+"()")
		}
		// it will be recognized as err if the sequence is different, sort before return
		sort.Strings(result)
		return result
	}
}
