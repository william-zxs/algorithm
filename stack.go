package main

import "fmt"

// 155. 最小栈
type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	//先判断最小值
	if val < this.GetMin() {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.GetMin())
	}
	//压栈
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {

	if len(this.stack) > 0 {
		//最小值 删除
		this.min = this.min[:len(this.min)-1]
		// 出栈
		this.stack = this.stack[:len(this.stack)-1]
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	return this.min[len(this.min)-1]
}

func main() {
	stack := Constructor()
	stack.Push(1)
	fmt.Println("==min==:", stack.GetMin())
	stack.Push(0)
	fmt.Println("==min==:", stack.GetMin())
	stack.Push(2)
	fmt.Println("==min==:", stack.GetMin())

	fmt.Println("==top==:", stack.Top())
	fmt.Println("==min==:", stack.GetMin())
	stack.Pop()
	fmt.Println("==min==:", stack.GetMin())
}
