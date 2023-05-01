package main

import "fmt"

func main() {
	fmt.Println(numOfMinutes(1, 0, []int{-1}, []int{0}))
}

type StaffNode struct {
	Val     int
	time    int
	members []*StaffNode
}

func createTree(n int, headID int, manager []int, informTime []int) *StaffNode {
	var head *StaffNode
	staffs := make([]*StaffNode, n)
	for i := 0; i < n; i++ {
		// create/update self node
		if staffs[i] == nil {
			staffs[i] = &StaffNode{
				Val:     i,
				members: []*StaffNode{},
			}
		}
		staffs[i].time = informTime[i]

		// if manager exists add this node to its staffs, otherwise create first
		if manager[i] == -1 {
			head = staffs[i]
		} else {
			if staffs[manager[i]] == nil {
				staffs[manager[i]] = &StaffNode{
					Val:     manager[i],
					members: []*StaffNode{},
				}
			}
			staffs[manager[i]].members = append(staffs[manager[i]].members, staffs[i])
		}

	}

	return head
}

func findLongest(head *StaffNode) int {
	longest := 0
	if len(head.members) > 0 {
		for _, sub := range head.members {
			longest = max(longest, findLongest(sub))
		}
	}
	return longest + head.time
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	head := createTree(n, headID, manager, informTime)
	return findLongest(head)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
