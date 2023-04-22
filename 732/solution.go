package main

import "fmt"

/*
*["MyCalendarThree","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book"]
[[],[47,50],[1,10],[27,36],[40,47],[20,27],[15,23],[10,18],[27,36],[17,25],[8,17],[24,33],[23,28],[21,27],
[47,50],[14,21],[26,32],[16,21],[2,7],[24,33],[6,13],[44,50],[33,39],[30,36],[6,15],[21,27],[49,50],[38,45],[4,12],[46,50],[13,21]]
*/
func main() {
	my := Constructor()
	fmt.Println(my.Book(47, 50))
	fmt.Println(my.Book(1, 10))
	fmt.Println(my.Book(27, 36))
	fmt.Println(my.Book(40, 47))
	fmt.Println(my.Book(20, 27))
	fmt.Println(my.Book(15, 23))
	fmt.Println(my.Book(10, 18))
	fmt.Println(my.Book(27, 36))
	fmt.Println(my.Book(17, 25))
	fmt.Println(my.Book(8, 17))
	fmt.Println(my.Book(24, 33))
	fmt.Println(my.Book(23, 28))
	fmt.Println(my.Book(21, 27))
	fmt.Println(my.Book(47, 50))
}

type MyCalendarThree struct {
	timeline map[int][]int
}

func Constructor() MyCalendarThree {
	mycal := MyCalendarThree{
		timeline: make(map[int][]int),
	}
	return mycal
}

func (m *MyCalendarThree) getTimelineKeys() []int {
	keys := make([]int, len(m.timeline))
	i := 0
	for k, _ := range m.timeline {
		keys[i] = k
		i++
	}
	//fmt.Println("keys", keys)
	return keys
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	//fmt.Println("========================")
	//fmt.Println("booking", startTime, endTime-1)
	max := 1
	leftMatch, rightMatch := endTime-1, startTime
	keys := this.getTimelineKeys()
	for _, k := range keys {
		v := this.timeline[k]
		// 四种场景
		// 1: k,v[0] 被包含于startTime, endTime
		// 2: k,v[0] 包含 startTime, endTime
		// 3: k,v[0] 被startTime截断
		// 4: k,v[0] 被endTime截断
		if k >= startTime && v[0] <= endTime-1 {
			// 1
			this.timeline[k] = []int{v[0], v[1] + 1}
			leftMatch = minNum(k-1, leftMatch)
			rightMatch = maxNum(v[0]+1, rightMatch)
		} else if k <= startTime && v[0] >= endTime-1 {
			// 2 分三段 [k, startTime-1] [startTime, endTime-1][endTime, v[0]]
			if k != startTime {
				this.timeline[k] = []int{startTime - 1, v[1]}
			}
			this.timeline[startTime] = []int{endTime - 1, v[1] + 1}
			if v[0] != endTime-1 {
				this.timeline[endTime] = []int{v[0], v[1]}
			}
			leftMatch = startTime - 1
			rightMatch = endTime
		} else if k < startTime && v[0] >= startTime {
			// 3 分两段 [k, startTime-1][startTime,v[0]]
			this.timeline[k] = []int{startTime - 1, v[1]}
			this.timeline[startTime] = []int{v[0], v[1] + 1}
			leftMatch = k
			rightMatch = maxNum(v[0]+1, rightMatch)
		} else if k <= endTime-1 && v[0] > endTime-1 {
			// 4 分两段 [k, endTime-1][endTime,v[0]]
			this.timeline[k] = []int{endTime - 1, v[1] + 1}
			this.timeline[endTime] = []int{v[0], v[1]}
			leftMatch = minNum(k-1, leftMatch)
			rightMatch = v[0]
		} else {
			max = maxNum(max, v[1])
			continue
		}
		max = maxNum(max, v[1]+1)
	}
	// 剩余未匹配到内容添加到时间线
	if leftMatch == endTime-1 {
		//fmt.Println("created", startTime, leftMatch)
		this.timeline[startTime] = []int{leftMatch, 1}
	} else {
		if leftMatch > startTime {
			//fmt.Println("created", startTime, leftMatch)
			this.timeline[startTime] = []int{leftMatch, 1}
		}
		if rightMatch < endTime {
			//fmt.Println("created", rightMatch, endTime-1)
			this.timeline[rightMatch] = []int{endTime - 1, 1}
		}
	}
	//fmt.Println(this.timeline)
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
