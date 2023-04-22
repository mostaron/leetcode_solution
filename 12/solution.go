package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func intToRoman(num int) string {
	dict := make(map[int]string)
	dict[1] = "I"
	dict[5] = "V"
	dict[10] = "X"
	dict[50] = "L"
	dict[100] = "C"
	dict[500] = "D"
	dict[1000] = "M"
	result := ""
	for i := 1; i <= num; i *= 10 {
		tmp := ""
		j := num / i % 10
		if j < 4 {
			for x := 0; x < j; x++ {
				tmp += dict[i]
			}
		} else if j == 4 {
			tmp = dict[i] + dict[5*i]
		} else if j == 5 {
			tmp = dict[i*5]
		} else if j < 9 {
			tmp = dict[i*5]
			for x := 0; x < j-5; x++ {
				tmp += dict[i]
			}
		} else if j == 9 {
			tmp = dict[i] + dict[i*10]
		}
		result = tmp + result

	}
	return result
}
