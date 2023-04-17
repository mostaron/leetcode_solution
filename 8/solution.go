package main

import "fmt"

func main() {
	fmt.Println(myAtoi(" -42"))
}

func myAtoi(s string) int {
	result := 0
	flag := 0
	for j := 0; j < len(s); j++ {
		c := s[j]
		if c >= '0' && c <= '9' {
			if flag == 0 {
				flag = 1
			}

			i := int(c) - int('0')

			result = result*10 + flag*i

			if result >= 1<<31-1 {
				return 1<<31 - 1
			}
			if result <= -1*1<<31 {
				return -1 * 1 << 31
			}

			continue
		}
		if flag == 0 && (c == '-' || c == '+' || c == ' ') {
			if c == '-' {
				flag = -1
			} else if c == '+' {
				flag = 1
			}
			continue
		}
		break

	}
	return result
}
