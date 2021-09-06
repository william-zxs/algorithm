package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//删除排序链表重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// 删除排序链表中的重复元素 II
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//  [1,2,3,3,4,4,5]
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	current := dummy

	for current != nil {

		rmNext := false
		for current.Next != nil && current.Next.Next != nil && current.Next.Val == current.Next.Next.Val {
			current.Next = current.Next.Next
			rmNext = true
		}

		if rmNext {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}

	}
	return dummy.Next
}

// 反转链表   递归的方式，还有些问题
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	_, headRev := doWork(head)
	return headRev
}

func doWork(node *ListNode) (*ListNode, *ListNode) {
	if node == nil {
		return nil, nil
	}

	res, head := doWork(node.Next)

	if res != nil {
		res.Next = node
	} else {
		head = node
	}
	fmt.Println("node==:", node)
	fmt.Println(node, " ", head)
	return node, head
}

//反转链表 循环的方式
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

// 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 原始：1->2->3->4->5->NULL  left=2  right=4
	// 目标：1->4->3->2->5->null

	var pre *ListNode
	dummyNode := &ListNode{
		Val: 0,
	}
	dummyNode.Next = head
	head = dummyNode

	var start *ListNode
	var startNext *ListNode
	i := 0

	for ; i < left; i++ {
		start = head
		head = head.Next
		fmt.Println("==i==:", i)
	}
	startNext = head
	for ; i <= right; i++ {
		fmt.Println("==i==:", i)
		nextTemp := head.Next
		head.Next = pre
		pre = head
		head = nextTemp
	}

	start.Next = pre
	startNext.Next = head

	return dummyNode.Next
}

func main() {

	//[1,2,3,4,5]
	node5 := ListNode{
		Val:  5,
		Next: nil,
	}

	node4 := ListNode{
		Val:  4,
		Next: &node5,
	}

	node3 := ListNode{
		Val:  3,
		Next: &node4,
	}

	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}
	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}

	head := reverseBetween(&node1, 1, 4)
	fmt.Println(head)
	for head != nil {
		fmt.Println("===node===:", head)
		head = head.Next
	}
}
