package main

import "fmt"

func main() {
	fmt.Println(convert("123456789", 3))
	fmt.Println(convert("123456789", 4))
	fmt.Println(convert("A", 1))
}

/*
*
the index of models should like follow
1       9
2     8 10
3   7   11
4 6     12
5       13

the calculation should be like this

	0					|                             | 1

0    0 * (2n-1) + 0		|                             | 1 * (2n-1) + 0
1    0 * (2n-1) + 1		| (x+1) * (2n-1) - 1          | 1 * (2n-1) + 1

	.........			| ......                      |
	.........			| ......                      |

n-2  0 * (2n-1) + (n-2)	| (x+1) * (2n-1) - (n-2)      | 1 * (2n-1) + (n-2)
n-1  0 * (2n-1) + n-1	|                             | 1* (2n-1) + n

the formula should like this (x=column, y=row)
0 x*2(n-1)
1 x*2(n-1)+y, (x+1)*2(n-1)-y
...
n-1 x*2(n-1)+(n-1) => (x*2+1)*(n-1)
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	cArray := []rune(s)
	result := make([]rune, len(cArray))
	//fmt.Println("len of cArray", len(cArray), "len of result", len(result))
	// first row (0) : x*(2n-1)+y
	pointer := 0
	for x := 0; x*2*(numRows-1) < len(cArray); x++ {
		result[pointer] = cArray[x*2*(numRows-1)]
		pointer++
	}
	//fmt.Println(string(result))
	// second to n-2 row
	for y := 1; y <= numRows-2; y++ {
		for x := 0; x*2*(numRows-1)+y < len(cArray); x++ {
			// left part
			result[pointer] = cArray[x*2*(numRows-1)+y]
			pointer++
			// right part
			if (x+1)*2*(numRows-1)-y < len(cArray) {
				result[pointer] = cArray[(x+1)*2*(numRows-1)-y]
				pointer++
			} else {
				// exceed the end, jump to next row
				break
			}
		}
		//fmt.Println(string(result))
	}
	// last row (n-1) : x*(2n-1)+y
	for x := 0; (x*2+1)*(numRows-1) < len(cArray); x++ {
		result[pointer] = cArray[(x*2+1)*(numRows-1)]
		pointer++
	}
	//fmt.Println(string(result))
	return string(result)
}
