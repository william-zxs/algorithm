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

type Node struct {
	Val       int
	Neighbors []*Node
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

//133. 克隆图
func cloneGraph(node *Node) *Node {
	//采用哈希表和递归  精髓在于用哈希表记录下已经clone的节点
	visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if v, ok := visited[node]; ok {
			return v
		}

		cloneNode := &Node{node.Val, []*Node{}}
		visited[node] = cloneNode

		for _, ne := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(ne))
		}
		return cloneNode
	}
	return cg(node)
}

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	//深度搜索优先  哈希表 记录下位置，位置是x、y 决定
	visited := map[string]int{}
	count := 0
	//遍历 二维切片
	for y := 0; y < len(grid); y++ {
		yline := grid[y]
		for x := 0; x < len(yline); x++ {
			if string(yline[x]) == "1" {
				//m的上下左右 四个方位都要搜索，如果找到"0"那么就是岛屿
				xStr := strconv.Itoa(x)
				yStr := strconv.Itoa(y)
				if _, ok := visited[xStr+yStr]; ok {
					continue
				}
				//肯定是一个岛屿
				count += 1

				//把连接的节点放进哈希表
				//右边
				for m := x + 1; m < len(yline); m++ {
					if string(yline[m]) == "0" {
						break
					} else {
						if _, ok := visited[strconv.Itoa(m)+yStr]; ok {
							continue
						} else {
							visited[strconv.Itoa(m)+yStr] = 1
						}

						//下边
						for n := y + 1; n < len(grid); n++ {
							if string(grid[n][m]) == "0" {
								break
							} else {
								if _, ok := visited[strconv.Itoa(m)+strconv.Itoa(n)]; ok {
									continue
								} else {
									visited[strconv.Itoa(m)+strconv.Itoa(n)] = 1
								}
							}
						}
					}
				}
				//下边
				for n := y + 1; n < len(grid); n++ {
					if string(grid[n][x]) == "0" {
						break
					} else {
						if _, ok := visited[xStr+strconv.Itoa(n)]; ok {
							continue
						} else {
							visited[xStr+strconv.Itoa(n)] = 1
						}
					}
					//右边
					for m := x + 1; m < len(yline); m++ {
						if string(grid[n][m]) == "0" {
							break
						} else {
							if _, ok := visited[strconv.Itoa(m)+strconv.Itoa(n)]; ok {
								continue
							} else {
								visited[strconv.Itoa(m)+strconv.Itoa(n)] = 1
							}
						}
					}
					//左边
					for m := x - 1; m >= 0; m-- {
						if string(grid[n][m]) == "0" {
							break
						} else {
							if _, ok := visited[strconv.Itoa(m)+strconv.Itoa(n)]; ok {
								continue
							} else {
								visited[strconv.Itoa(m)+strconv.Itoa(n)] = 1
							}
						}
					}
				}
			}
		}
	}
	//统计数量
	return count
}

func numIslands2(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	count := 0
	xLength := len(grid[0])
	yLength := len(grid)
	var dfsIslands func(y, x int)
	dfsIslands = func(y, x int) {
		if !(y >= 0 && y < yLength && x >= 0 && x < xLength) {
			return
		}

		if string(grid[y][x]) == "1" {
			grid[y][x] = 0
			//上
			dfsIslands(y-1, x)
			//下
			dfsIslands(y+1, x)
			//左
			dfsIslands(y, x-1)
			//右
			dfsIslands(y, x+1)
		}
	}

	for y := 0; y < yLength; y++ {
		for x := 0; x < xLength; x++ {
			if string(grid[y][x]) == "1" {
				count += 1
				dfsIslands(y, x)
			}
		}
	}
	return count
}

// 84. 柱状图中最大的矩形
func largestRectangleArea2(heights []int) int {
	// 此种方式需要两层循环，时间复杂度是O(n^2),在leedcode会超时
	maxArea := 0
	for cur := 0; cur < len(heights); cur++ {
		// stack := make([]int, 0)
		minStack := make([]int, 0)
		for i := cur; i < len(heights); i++ {
			if len(minStack) == 0 {
				minStack = append(minStack, heights[i])
			} else {
				if minStack[len(minStack)-1] > heights[i] {
					minStack = append(minStack, heights[i])
				} else {
					minStack = append(minStack, minStack[len(minStack)-1])
				}
			}
			area := len(minStack) * minStack[len(minStack)-1]
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// 84. 柱状图中最大的矩形
func largestRectangleArea3(heights []int) int {
	//一层循环的方式   这个效率不高，也会超时
	h := 0
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		h = heights[i]
		lcount := 0
		rcount := 0
		//等于0的情况不用考虑
		if h > 0 {
			// left
			for m := i - 1; m >= 0; m-- {
				hl := heights[m]
				if hl < h {
					break
				}
				lcount++
			}
			// right
			for m := i + 1; m < len(heights); m++ {
				hr := heights[m]
				if hr < h {
					break
				}
				rcount++
			}

			area := (lcount + rcount + 1) * h
			if area > maxArea {
				maxArea = area
			}
		}

	}
	return maxArea
}

// 84. 柱状图中最大的矩形
// 单调栈(mono_stack)问题
func largestRectangleArea(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))
	monoStack := make([]int, 0)
	maxArea := 0

	// 计算left
	for i := 0; i < len(heights); i++ {

		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}

	// 计算right
	monoStack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
		fmt.Println("==right==", right)
	}
	// left [-1,0]
	// right [2,2]
	//       2*0  1*9

	fmt.Println("==right==", right, "==left==", left)
	for i := 0; i < len(heights); i++ {
		area := heights[i] * (right[i] - left[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//单调栈  42、 84、739、496、316、901、402、581 题
//我们用一个具体的例子 [6,7,5,2,4,5,9,3] 来帮助读者理解单调栈。
//我们需要求出每一根柱子的左右两侧且最近的小于其高度的柱子。初始时的栈为空。
func monotonic() {
	data := []int{6, 7, 5, 2, 4, 5, 9, 3}
	// 1 4 3 2 4
	indexStack := make([]int, len(data))
	monoStack := make([]int, 0)
	//左边
	for i := 0; i < len(data); i++ {
		count := 1
		fmt.Println("==monoStack==", monoStack)
		for len(monoStack) > 0 && monoStack[len(monoStack)-1] >= data[i] {
			monoStack = monoStack[:len(monoStack)-1]
			count++
		}

		if len(monoStack) == 0 {
			indexStack[i] = -1
		} else {
			indexStack[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, data[i])
	}
	fmt.Print("==indexStack==", indexStack)
}

func main() {

	// data := [][]byte{
	// 	{[]byte("1")[0], []byte("0")[0], []byte("0")[0], []byte("1")[0], []byte("0")[0]},
	// 	{[]byte("0")[0], []byte("1")[0], []byte("0")[0], []byte("1")[0], []byte("0")[0]}}

	//[["1","1","1"],
	//["0","1","0"],
	//["1","1","1"]]

	// 深度搜索优先

	// res := numIslands2(data)
	// data := []int{2, 1, 5, 6, 2, 3}
	// res := largestRectangleArea(data)
	// fmt.Println("==res==", res)
	data := []int{0, 9}
	res := largestRectangleArea(data)
	fmt.Println("==res==", res)
}
