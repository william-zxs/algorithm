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

func main() {

	node2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	head := node1

	res := swapPairs2(head)
	fmt.Println("==res==:", res)
}
