package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

//394. 字符串解码
func decodeString(s string) string {
	// 用栈的方式
	stack := make([]byte, 0)
	for _, item := range s {
		if string(item) == "]" {
			// 找到上一个"["
			tempStack := make([]byte, 0)
			for i := len(stack) - 1; i >= 0; i-- {
				if string(stack[i]) == "[" {
					stack = stack[:i]
					break
				}
				tempStack = append([]byte{stack[i]}, tempStack...)
			}

			countStack := make([]byte, 0)
			for i := len(stack) - 1; i >= 0; i-- {
				_, err := strconv.Atoi(string(stack[i]))
				if err != nil {
					if len(countStack) == 0 {
						return "==err=="
					}
					break
				}
				countStack = append([]byte{stack[i]}, countStack...)
				stack = stack[:len(stack)-1]
			}

			count, err := strconv.Atoi(string(countStack))
			if err != nil {
				return "==err=="
			}
			tempCopy := tempStack
			for i := 0; i < count-1; i++ {
				tempStack = append(tempStack, tempCopy...)
			}
			stack = append(stack, tempStack...)
		} else {
			stack = append(stack, byte(item))
		}

	}

	return string(stack)
}

//394. 字符串解码
func decodeString2(s string) string {
	// 递归的方式
	// "e3[a10[b]]c2[dd]"
	data, err, _ := parseStr(s)
	if err != "" {
		return err
	}
	return string(data)
}

//按照倍数构造切片
func multiCount(stack []byte, strStack []byte) ([]byte, string) {
	countStack := make([]byte, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(string(stack[i]))
		if err != nil {
			if len(countStack) == 0 {
				fmt.Println("==stack[i]==", string(stack[i]))
				return nil, "==strconv.Atoi==err1"
			}
			break
		}
		countStack = append([]byte{(stack)[i]}, countStack...)
		stack = stack[:len(stack)-1]
	}

	count, err := strconv.Atoi(string(countStack))
	if err != nil {
		return nil, "==strconv.Atoi==err2"
	}
	for i := 0; i < count; i++ {
		stack = append(stack, strStack...)
	}
	fmt.Println("==multiCount==", string(stack))
	return stack, ""
}

// a10[b]
func parseStr(s string) ([]byte, string, string) {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "[" {
			strStack, err, resStr := parseStr(s[i+1:])
			if err != "" {
				return nil, err, resStr
			}
			fmt.Println("==strStack==", string(strStack))
			stack, err = multiCount(stack, strStack)
			if resStr != "" {
				resStrStack, err, resStr := parseStr(resStr)
				return append(stack, resStrStack...), err, resStr
			}
			return stack, "", ""

		} else if string(s[i]) == "]" {
			if len(s)-1 > i {
				return stack, "", s[i+1:]
			}
			return stack, "", ""
		} else if s[i] >= '0' && s[i] <= '9' {
			//数字
			stack = append(stack, s[i])
		} else {
			//字母
			stack = append(stack, s[i])
		}
	}
	if len(stack) == 0 {
		return nil, "==1==", ""
	}
	return stack, "", ""
}

//94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) (list []int) {
	// 栈的方式
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		list = append(list, node.Val)
		root = node.Right
	}
	return
}

//递归的方式
func inorderTraversal2(root *TreeNode) (list []int) {
	return inorderWork(root)
}

func inorderWork(root *TreeNode) (list []int) {
	if root == nil {
		return
	}
	resLeft := inorderWork(root.Left)
	list = append(list, resLeft...)
	list = append(list, root.Val)
	resRight := inorderWork(root.Right)
	list = append(list, resRight...)
	return
}

func main() {

	// data := []string{"4", "13", "5", "/", "+"}
	// res := evalRPN(data)
	// byte() 怎么用的
	//abbbbabbbbabbbb ccdddddddddd
	s := "3[a2[bb]]cc10[d]"
	res := decodeString2(s)
	fmt.Println("==res==", res)
}
