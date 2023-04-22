package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	fmt.Println(countDaysTogether("08-06", "12-08", "02-04", "09-01"))
}

var DAY_OF_MONTH = []byte{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func convertToMD(str string) []byte {
	var result []byte
	for _, s := range strings.Split(str, "-") {
		i, err := strconv.Atoi(s)
		if err != nil {
			errors.New("format error")
			break
		}
		result = append(result, byte(i))
	}
	return result
}

func isGreater(a *[]byte, b *[]byte) bool {
	return (*a)[0] > (*b)[0] || ((*a)[0] == (*b)[0] && (*a)[1] > (*b)[1])
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var aStart, aEnd, bStart, bEnd, pStart, pEnd []byte
	aStart = convertToMD(arriveAlice)
	aEnd = convertToMD(leaveAlice)
	bStart = convertToMD(arriveBob)
	bEnd = convertToMD(leaveBob)

	if isGreater(&aStart, &bStart) {
		pStart = aStart
	} else {
		pStart = bStart
	}
	if isGreater(&aEnd, &bEnd) {
		pEnd = bEnd
	} else {
		pEnd = aEnd
	}

	if isGreater(&pStart, &pEnd) {
		return 0
	}

	fmt.Println(pStart, pEnd)
	dayCount := int(pEnd[1]) - int(pStart[1]) + 1
	for i := pStart[0]; i < pEnd[0]; i++ {
		dayCount += int(DAY_OF_MONTH[i-1])
	}

	return dayCount

}
