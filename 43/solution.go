package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("66", "66"))
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := ""
	tmp := make([]uint8, len(num1)+len(num2))
	for i := 1; i <= len(num1); i++ {
		for j := 1; j <= len(num2); j++ {
			value := (num1[len(num1)-i] - '0') * (num2[len(num2)-j] - '0')
			start := i + j - 1
			value += tmp[len(tmp)-start]
			overflow := value / 10
			tmp[len(tmp)-start] = value % 10
			for overflow != 0 {
				start++
				value = tmp[len(tmp)-start] + overflow
				tmp[len(tmp)-start] = value % 10
				overflow = value / 10
			}
		}
	}
	headZero := true
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 0 && headZero {
			continue
		}
		headZero = false
		result += string(tmp[i] + '0')
	}
	return result
}
func add(num1, num2 string) string {
	result := ""
	overflow := uint8(0)
	for i := 1; i <= len(num1) || i <= len(num2); i++ {
		bit := uint8(0)
		if i <= len(num1) {
			bit += num1[len(num1)-i] - '0'
		}
		if i <= len(num2) {
			bit += num2[len(num2)-i] - '0'
		}
		bit += overflow
		result = string(bit%10+'0') + result
		overflow = bit / uint8(10)
	}
	if overflow != uint8(0) {
		result = string(overflow+'0') + result
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
