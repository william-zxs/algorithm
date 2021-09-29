package main

/*
队列
常用于 BFS 宽度优先搜索
*/
//232. 用栈实现队列  fifo
type MyQueue struct {
	stack []int
}

func Constructor2() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	data := this.stack[0]
	this.stack = this.stack[1:]
	return data
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	if count := len(this.stack); count > 0 {
		return false
	} else {
		return true
	}
}

func main() {

}
