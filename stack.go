package main

import (
	"fmt"
	"strconv"
)

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

// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	//["4","13","5","/","+"]
	// 符号  + - * /
	stack := make([]int, 0)

	for _, token := range tokens {

		switch token {
		// 	golang 的 switch 是不下坠的， java 的 switch是会下坠的
		case "+", "-", "*", "/":
			intOne := stack[len(stack)-2]
			intTwo := stack[len(stack)-1]

			var res int
			switch token {
			case "+":
				res = intOne + intTwo
			case "-":
				res = intOne - intTwo
			case "*":
				res = intOne * intTwo
			case "/":
				res = intOne / intTwo
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		default:
			intData, _ := strconv.Atoi(token)
			stack = append(stack, intData)
		}

	}
	return stack[0]
}

func main() {

	data := []string{"4", "13", "5", "/", "+"}
	// res := evalRPN(data)

	fmt.Println("==res==", data)
}
