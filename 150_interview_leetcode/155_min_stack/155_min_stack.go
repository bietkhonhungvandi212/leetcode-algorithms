package minstack

import "math"

type Node struct {
	val      int
	localMin int
}

type MinStack struct {
	min   int
	stack []Node
}

func Constructor() MinStack {
	min := math.MaxInt
	stack := []Node{}

	return MinStack{min, stack}
}

func (this *MinStack) Push(val int) {

	if val < this.min {
		this.min = val
	}

	node := Node{val: val, localMin: this.min}

	this.stack = append(this.stack, node)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	this.stack = this.stack[:len(this.stack)-1]

	if len(this.stack) == 0 {
		this.min = math.MaxInt
	}

	this.min = this.stack[len(this.stack)-1].localMin
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
