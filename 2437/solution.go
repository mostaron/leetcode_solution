package main

import "fmt"

func main() {
	fmt.Println(countTime("?5:00"))
	fmt.Println(countTime("0?:0?"))
	fmt.Println(countTime("??:??"))
}

func countTime(time string) int {
	hours, minutes := 1, 1
	switch {
	case time[0] == '?' && time[1] == '?':
		hours = 24
	case time[0] != '?' && time[1] != '?':
		hours = 1
	case time[0] == '0' || time[0] == '1':
		hours = 10
	case time[0] == '2':
		hours = 4
	case time[1] < '4':
		hours = 3
	case time[1] >= '4':
		hours = 2
	}
	switch {
	case time[3] == '?' && time[4] == '?':
		minutes = 60
	case time[3] != '?' && time[4] != '?':
		minutes = 1
	case time[3] != '?':
		minutes = 10
	case time[4] != '?':
		minutes = 6
	}
	return hours * minutes
}
