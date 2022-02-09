package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
二刷
*/

/*
141
142
138
394. 字符串解码

84. 柱状图中最大的矩形 困难

542. 01 矩阵
136. 只出现一次的数字
137. 只出现一次的数字 II
260. 只出现一次的数字 III
191. 位1的个数
338. 比特位计数
190. 颠倒二进制位
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 28. 实现 strStr()
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

//78. 子集
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	list := make([]int, 0)
	subsetsHelper(0, nums, list, &result)
	return result
}

func subsetsHelper(pos int, nums []int, list []int, result *[][]int) {

	data := make([]int, len(list))
	copy(data, list)
	*result = append(*result, data)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		subsetsHelper(i+1, nums, list, result)
		list = list[:len(list)-1]
	}
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthHelper(root)
}

func maxDepthHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthHelper(root.Left)
	right := maxDepthHelper(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	//分治法
	if root == nil {
		return true
	}
	_, b := isBalancedHelper(root)
	return b
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftD, leftB := isBalancedHelper(root.Left)
	rightD, rightB := isBalancedHelper(root.Right)
	if !leftB || !rightB {
		return 0, false
	}
	if leftD-rightD > 1 || rightD-leftD > 1 {
		return 0, false
	}

	if leftD > rightD {
		return leftD + 1, true
	}
	return rightD + 1, true

}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, node := helper(root, p, q, 0)
	return node
}

func helper(root, p, q *TreeNode, count int) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	left, leftNode := helper(root.Left, p, q, count)
	right, rightNode := helper(root.Right, p, q, count)
	if left == 2 {
		return 2, leftNode
	}
	if right == 2 {
		return 2, rightNode
	}

	count = 0
	if root == p || root == q {
		count += 1
	}
	if left == 1 {
		count += 1
	}
	if right == 1 {
		count += 1
	}
	if count == 2 {
		return count, root
	}
	return count, nil
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	//贡献值的做法
	maxSum := root.Val

	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftGain := max(maxGain(root.Left), 0)
		rightGain := max(maxGain(root.Right), 0)

		maxSum = max(leftGain+rightGain+root.Val, maxSum)

		return max(leftGain, rightGain) + root.Val
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {

	levelList := make([]*TreeNode, 0)
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	levelList = append(levelList, root)

	for len(levelList) > 0 {
		data := make([]int, 0)
		l := len(levelList)
		for i := 0; i < l; i++ {
			node := levelList[i]
			data = append(data, node.Val)
			if node.Left != nil {
				levelList = append(levelList, node.Left)
			}
			if node.Right != nil {
				levelList = append(levelList, node.Right)
			}
		}
		result = append(result, data)
		levelList = levelList[l:]
	}
	return result
}

// 107. 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {

	result := make([][]int, 0)
	levelStack := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	levelStack = append(levelStack, root)
	for len(levelStack) > 0 {
		l := len(levelStack)
		data := make([]int, 0)
		for i := 0; i < l; i++ {
			node := levelStack[i]
			data = append(data, node.Val)
			if node.Left != nil {
				levelStack = append(levelStack, node.Left)
			}
			if node.Right != nil {
				levelStack = append(levelStack, node.Right)
			}
		}
		levelStack = levelStack[l:]
		result = append(result, data)
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

//103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	i := 0
	for len(stack) > 0 {
		l := len(stack)
		data := make([]int, 0)
		for i := 0; i < l; i++ {
			node := stack[i]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			data = append(data, node.Val)
		}

		if i%2 != 0 {
			reverseSlice(data)
		}
		result = append(result, data)
		stack = stack[l:]
		i++
	}
	return result
}

func reverseSlice(data []int) []int {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

// 98. 验证二叉搜索树
/*
二叉搜索树  的中序遍历是递增的
还可以用分治法。
*/
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			node = node.Right
			for node != nil {
				stack = append(stack, node)
				node = node.Left
			}
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

//DFS 分治法
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isTrue, _, _ := doWorkIsValidBST2(root)
	if isTrue {
		return true
	}
	return false
}

func doWorkIsValidBST2(root *TreeNode) (bool, int, int) {
	if root.Left != nil && root.Right != nil {
		isLeft, leftMin, leftMax := doWorkIsValidBST2(root.Left)
		isRight, rightMin, rightMax := doWorkIsValidBST2(root.Right)
		if isLeft && isRight && root.Val > leftMax && root.Val < rightMin {
			return true, leftMin, rightMax
		} else {
			return false, 0, 0
		}
	}

	if root.Left != nil {
		isLeft, leftMin, leftMax := doWorkIsValidBST2(root.Left)
		if isLeft && root.Val > leftMax {
			return true, leftMin, root.Val
		} else {
			return false, 0, 0
		}
	}
	if root.Right != nil {
		isRight, rightMin, rightMax := doWorkIsValidBST2(root.Right)
		if isRight && root.Val < rightMin {
			return true, root.Val, rightMax
		} else {
			return false, 0, 0
		}
	}

	return true, root.Val, root.Val

}

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	head := root
	for root != nil {
		if root.Val < val {
			if root.Right == nil {
				root.Right = &TreeNode{Val: val}
				return head
			}
			root = root.Right
		}
		if root.Val > val {
			if root.Left == nil {
				root.Left = &TreeNode{Val: val}
				return head
			}
			root = root.Left
		}
	}

	return head
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	headR := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}

	}
	return headR

}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: -1000}
	dummy.Next = head
	node := dummy
	dupMap := make(map[int]bool, 0)

	for dummy.Next != nil {
		if _, ok := dupMap[dummy.Next.Val]; ok {
			dummy.Next = dummy.Next.Next
			continue
		}

		if dummy.Next.Next != nil {
			if dummy.Next.Val == dummy.Next.Next.Val {
				dupMap[dummy.Next.Val] = true
				dummy.Next = dummy.Next.Next
				continue
			}
		}
		dummy = dummy.Next

	}

	return node.Next
}

// 92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var pre *ListNode
	i := 0

	var firstHead *ListNode
	var secondHead *ListNode
	var thirdHead *ListNode
	var finalHead *ListNode

	firstHead = head
	for head != nil {
		i++
		if i == left {
			secondHead = head
			if pre == nil {
				firstHead = nil
			} else {
				pre.Next = nil
			}
		}
		if i == right {
			if head.Next == nil {
				thirdHead = nil
			} else {
				thirdHead = head.Next
			}
			head.Next = nil
		}
		pre = head
		head = head.Next
	}
	reSecondHead := reverseList(secondHead)

	if firstHead == nil {
		finalHead = reSecondHead
	} else {
		finalHead = firstHead
		for firstHead != nil && firstHead.Next != nil {
			firstHead = firstHead.Next
		}
		firstHead.Next = reSecondHead
	}

	for reSecondHead != nil && reSecondHead.Next != nil {
		reSecondHead = reSecondHead.Next
	}
	reSecondHead.Next = thirdHead
	return finalHead
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre

		if next == nil {
			dummy.Next = head
			break
		}
		pre = head
		head = next
	}
	return dummy.Next
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	resHead := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}

	if list1 == nil {
		dummy.Next = list2
	} else {
		dummy.Next = list1
	}
	return resHead.Next
}

// 86. 分隔链表
func partition(head *ListNode, x int) *ListNode {
	left := &ListNode{Val: 0}
	right := &ListNode{Val: 0}
	rightHead := right
	leftHead := left
	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = next
	}
	left.Next = rightHead.Next
	return leftHead.Next
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

//归并
func mergeSort(head *ListNode) *ListNode {
	//判断终止条件
	if head == nil || head.Next == nil {
		return head
	}
	//找到中点
	rightHead := findMiddle(head)
	leftHead := head
	//分两部分递归
	leftOrderHead := mergeSort(leftHead)
	rightOrderHead := mergeSort(rightHead)
	//将两个有序链表合并为一个有序链表
	head = mergeNodeList(leftOrderHead, rightOrderHead)
	//返回该链表
	return head
}

func findMiddle(head *ListNode) *ListNode {
	//快慢指针的方式找到中间的node
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	rightHead := slow.Next
	slow.Next = nil
	return rightHead
}

//148. 排序链表
/*
	使用归并排序来计算，其实是一个排序，其他排序也都可以
*/
func mergeNodeList(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	dummy := &ListNode{Val: 0}
	head := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			dummy.Next = left
			left = left.Next
		} else {
			dummy.Next = right
			right = right.Next
		}
		dummy = dummy.Next
	}

	if left == nil {
		dummy.Next = right
	}
	if right == nil {
		dummy.Next = left
	}
	return head.Next
}

// 143. 重排链表
func reorderList(head *ListNode) {
	//分成两部分
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	right := slow.Next
	left := head
	slow.Next = nil
	//第二部分反转
	var pre *ListNode
	var reRight *ListNode
	for right != nil {
		next := right.Next
		if next == nil {
			reRight = right
		}
		right.Next = pre
		pre = right
		right = next

	}
	//合并两部分
	for left != nil && reRight != nil {
		leftNext := left.Next
		rightNext := reRight.Next
		left.Next = reRight
		reRight.Next = leftNext
		left = leftNext
		reRight = rightNext
	}
}

// 141. 环形链表 *
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	s := head
	f := head

	s = s.Next
	f = f.Next.Next
	for s != f {
		if f == nil || f.Next == nil {
			return nil
		}
		s = s.Next
		f = f.Next.Next
	}
	f = head
	for s != f {
		s = s.Next
		f = f.Next
	}
	return s

}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	for i, j := 0, len(list)-1; i < len(list)/2; i, j = i+1, j-1 {
		if list[i] != list[j] {
			return false
		}
	}
	return true
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 138. 复制带随机指针的链表
func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val, Next: next}
		cur = next
	}

	cur = head
	for cur != nil {
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	copyHead := head.Next
	cur = head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = cur.Next.Next
		cur = next
	}
	return copyHead
}

/*
栈和队列
栈的特点是后入先出
栈常用于DFS
队列常用于BFS
*/

// 155. 最小栈
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {

	return MinStack{stack: make([]int, 0), minStack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {

		if val < this.minStack[len(this.minStack)-1] {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
		}
	}

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		item := tokens[i]
		num, err := strconv.Atoi(item)
		if err == nil {
			stack = append(stack, num)
		} else {
			var val int
			switch item {
			case "+":
				val = stack[len(stack)-1] + stack[len(stack)-2]
			case "-":
				val = stack[len(stack)-2] - stack[len(stack)-1]
			case "*":
				val = stack[len(stack)-2] * stack[len(stack)-1]
			case "/":
				val = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// 394. 字符串解码
func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0)

	var subStr string
	var subNum int
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			numStack = append(numStack, subNum)
			strStack = append(strStack, subStr)
			subStr = ""
			subNum = 0
		} else if s[i] == ']' {
			popNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			popStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			repeatSubS := strings.Repeat(subStr, popNum)
			subStr = popStr + repeatSubS
		} else if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			subNum = subNum*10 + num
		} else {
			subStr += string(s[i])
		}

	}

	return subStr
}

// 94. 二叉树的中序遍历  * 基础  栈和递归的方式都要熟练，官方的答案比较简洁
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			subNode := node.Right
			for subNode != nil {
				stack = append(stack, subNode)
				subNode = subNode.Left
			}
		}
	}
	return result
}

type Node struct {
	Val       int
	Neighbors []*Node
}

// 133. 克隆图
func cloneGraph(node *Node) *Node {
	//BFS的方式

	if node == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node, 0)
	nodeMap[node] = &Node{Val: node.Val}
	stack := []*Node{node}
	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]
		for _, neighbor := range n.Neighbors {
			if _, ok := nodeMap[neighbor]; !ok {
				nodeMap[neighbor] = &Node{Val: neighbor.Val}
				stack = append(stack, neighbor)
			}
			nodeMap[n].Neighbors = append(nodeMap[n].Neighbors, nodeMap[neighbor])
		}
	}
	return nodeMap[node]
}

func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	copyMap := make(map[int]*Node, 0)
	var helper func(node *Node)
	helper = func(node *Node) {
		if _, ok := copyMap[node.Val]; !ok {
			copyNode := &Node{Val: node.Val}
			copyMap[node.Val] = copyNode
			neigList := make([]*Node, 0)
			if len(node.Neighbors) > 0 {
				for i := 0; i < len(node.Neighbors); i++ {
					subNode := node.Neighbors[i]
					helper(subNode)
					neigList = append(neigList, copyMap[subNode.Val])
				}
			}
			copyNode.Neighbors = neigList
		}
	}
	helper(node)
	return copyMap[node.Val]
}

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				helper(grid, i, j)
			}

		}
	}
	return count
}

func helper(grid [][]byte, i int, j int) {
	if !(i < len(grid) && i >= 0) || !(j < len(grid[0]) && j >= 0) {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '2'
		// 右边
		helper(grid, i, j+1)
		// 下边
		helper(grid, i+1, j)
		// 左边
		helper(grid, i, j-1)
		// 上边
		helper(grid, i-1, j)
	}
}

// 84. 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	if heights == nil {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	stack := []int{0}
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	maxArea := 0
	for i := 1; i < len(heights); i++ {
		for heights[i] < heights[stack[len(stack)-1]] {

			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			area := height * width
			maxArea = max(area, maxArea)
		}
		stack = append(stack, i)
	}
	return maxArea
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 232. 用栈实现队列
type MyQueue struct {
	Data []int
}

func Constructor() MyQueue {
	return MyQueue{Data: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.Data = append(this.Data, x)
}

func (this *MyQueue) Pop() int {
	num := this.Data[0]
	this.Data = this.Data[1:len(this.Data)]
	return num
}

func (this *MyQueue) Peek() int {
	num := this.Data[0]
	return num
}

func (this *MyQueue) Empty() bool {
	if len(this.Data) > 0 {
		return false
	}
	return true
}

// 542. 01 矩阵
func updateMatrix(mat [][]int) [][]int {
	//BFS 一般会用到栈
	stack := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				stack = append(stack, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	xl := len(mat)
	yl := len(mat[0])
	mx := []int{-1, 1, 0, 0}
	my := []int{0, 0, -1, 1}
	for len(stack) > 0 {
		point := stack[0]
		stack = stack[1:]
		for i := 0; i < 4; i++ {
			x := point[0] + mx[i]
			y := point[1] + my[i]
			if x >= 0 && x < xl && y >= 0 && y < yl && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				stack = append(stack, []int{x, y})
			}
		}
	}
	return mat
}

/*
二进制
*/

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	//异或
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

// 137. 只出现一次的数字 II
func singleNumberII(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		sumPosition := int32(0)
		for _, num := range nums {
			sumPosition += int32(num) >> i & 1
		}
		res = res | ((sumPosition % 3) << i)
	}
	return int(res)
}

// 260. 只出现一次的数字 III
func singleNumberIII(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	splitNum := xorSum & -xorSum
	lNum, rNum := 0, 0
	for _, num := range nums {
		if num&splitNum > 0 {
			lNum ^= num
		} else {
			rNum ^= num
		}
	}
	return []int{lNum, rNum}
}

//191. 位1的个数
func hammingWeight(num uint32) (ones int) {
	//n & (n−1)，其运算结果恰为把 nn 的二进制位中的最低位的 11 变为 00 之后的结果
	for ; num > 0; num &= num - 1 {
		ones++
		fmt.Println("num==", num)
	}
	return
}

//338. 比特位计数
func countBits(n int) []int {
	// dp的思想
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

//190. 颠倒二进制位
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num = num >> 1
	}
	return result
}
