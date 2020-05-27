package leetcode

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	arr    []int
	minArr []int
	pos    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		arr:    []int{},
		minArr: []int{},
		pos:    -1,
	}
}

func (this *MinStack) Push(x int) {
	var min int
	if this.pos == -1 {
		min = x
	} else {
		min = this.GetMin()
		if x < min {
			min = x
		}
	}

	this.arr = append(this.arr, x)
	this.minArr = append(this.minArr, min)
	this.pos++
	return
}

func (this *MinStack) Pop() {

	this.arr = this.arr[:this.pos]
	this.minArr = this.minArr[:this.pos]

	this.pos--
	return
}

func (this *MinStack) Top() int {
	if this.pos < 0 {
		return -1
	}
	return this.arr[this.pos]
}

func (this *MinStack) GetMin() int {
	if this.pos < 0 {
		return 0
	}
	return this.minArr[this.pos]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
