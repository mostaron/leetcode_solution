package main

func main() {

}

type MyCircularDeque struct {
	cap         int
	size        int
	value       int
	front, last *Node
}
type Node struct {
	value       int
	front, back *Node
}

func Constructor(k int) MyCircularDeque {
	m := MyCircularDeque{
		cap:  k,
		size: 0,
	}
	return m
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size == this.cap {
		return false
	}
	node := &Node{
		value: value,
	}
	if this.front == nil {
		this.last = node
	} else {
		this.front.front = node
		node.back = this.front
	}
	this.front = node
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size == this.cap {
		return false
	}
	node := &Node{
		value: value,
	}
	if this.front == nil {
		this.front = node
	} else {
		this.last.back = node
		node.front = this.last
	}
	this.last = node
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.size == 0 {
		return false
	}
	removed := this.front
	if removed.back == nil {
		this.front = nil
		this.last = nil
	} else {
		this.front = this.front.back
		this.front.front = nil
	}
	this.size--
	removed.back = nil
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}
	removed := this.last
	if removed.front == nil {
		this.front = nil
		this.last = nil
	} else {
		this.last = this.last.front
		this.last.back = nil
	}
	this.size--
	removed.front = nil
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}
	return this.front.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}
	return this.last.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.cap
}
