package main

func main() {

}

type MyCalendar struct {
	events [][]int
}

func Constructor() MyCalendar {
	mc := MyCalendar{
		events: [][]int{},
	}
	return mc
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, event := range this.events {
		if event[0] < end && event[1] > start {
			return false
		}
	}
	this.events = append(this.events, []int{start, end})
	return true
}
