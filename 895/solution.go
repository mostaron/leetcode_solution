package main

func main() {
	fs := Constructor()
	fs.Push(4)
	fs.Push(0)
	fs.Push(9)
	fs.Push(3)
	fs.Push(4)
	fs.Push(2)
	fs.Pop()
	fs.Push(6)
	fs.Pop()
	fs.Push(1)
	fs.Pop()
	fs.Push(1)
	fs.Pop()
	fs.Push(4)
	fs.Pop()
	fs.Pop()
	fs.Pop()
	fs.Pop()
	fs.Pop()
	fs.Pop()
}

type FreqStack struct {
	stack   map[int][]int
	freqs   map[int]int
	maxFreq int
}

func Constructor() FreqStack {
	fs := FreqStack{
		stack:   make(map[int][]int),
		freqs:   make(map[int]int),
		maxFreq: 0,
	}
	return fs
}

func (this *FreqStack) Push(val int) {
	this.freqs[val]++
	freq := this.freqs[val]
	if freq > this.maxFreq {
		this.maxFreq = freq
	}
	this.stack[freq] = append(this.stack[freq], val)
}

func (this *FreqStack) Pop() int {
	maxFreqArr := this.stack[this.maxFreq]
	removed := maxFreqArr[len(maxFreqArr)-1]
	maxFreqArr = maxFreqArr[0 : len(maxFreqArr)-1]
	if len(maxFreqArr) == 0 {
		delete(this.stack, this.maxFreq)
		this.maxFreq--
	} else {
		this.stack[this.maxFreq] = maxFreqArr
	}
	this.freqs[removed]--
	return removed
}
