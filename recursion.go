package main

import "fmt"

// 344. 反转字符串
func reverseString(s []byte) {
	// 1 2 3 4 5   2
	// 1 2 3 4     2
	l := len(s)
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	fmt.Println("==s==", s)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	// 循环迭代的方式
	var pre *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := head.Next
	for head != nil && head.Next != nil {

		next := head.Next
		temp := next.Next
		head.Next = temp
		next.Next = head
		if pre != nil {
			pre.Next = next
		}
		pre = head
		head = temp
	}
	return dummyHead
}

// 24. 两两交换链表中的节点
func swapPairs2(head *ListNode) *ListNode {
	// 递归的方式
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	next := head.Next
	next.Next = head
	head.Next = swapPairs2(nextHead)
	return next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	//二叉搜索树的特点是  左节点值小于根节点，右节点值大于根节点,左子树和右子树同样也是二叉搜索树

	if n == 0 {
		return nil
	}
	return generate(1, n)
}
func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	treeNodes := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {

		leftNodes := generate(start, i-1)
		rightNodes := generate(i+1, end)
		for j := 0; j < len(leftNodes); j++ {
			for k := 0; k < len(rightNodes); k++ {

				root := &TreeNode{Val: i}
				root.Left = leftNodes[j]
				root.Right = rightNodes[k]
				treeNodes = append(treeNodes, root)
				fmt.Println("==treeNodes==", treeNodes)
			}
		}
	}
	return treeNodes
}

// 509. 斐波那契数

func fib(n int) int {

	var help func(n int) int
	cache := make(map[int]int)

	help = func(n int) int {
		if num, ok := cache[n]; ok {
			return num
		}
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		res := help(n-1) + help(n-2)
		cache[n] = res
		return res
	}
	return help(n)
}

func main() {

	// node2 := &ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }
	// node1 := &ListNode{
	// 	Val:  1,
	// 	Next: node2,
	// }
	// head := node1

	// res := swapPairs2(head)

	// res := generateTrees(3)
	res := fib(3)
	fmt.Println("==res==:", res)
}
