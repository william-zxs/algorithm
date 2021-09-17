package main

import (
	"fmt"
)

//链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//根据切片构造链表
func buildLinkedList(list []int) *ListNode {
	dummyNode := &ListNode{
		Val: 0,
	}
	head := dummyNode
	for _, num := range list {
		fmt.Println("==num==:", num)
		node := &ListNode{
			Val: num,
		}
		head.Next = node
		head = node
	}
	return dummyNode.Next
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
	dummyNode := &ListNode{
		Val: 0,
	}
	dummyHead := dummyNode
	_ = reverseLinkedList(head, dummyHead)
	// fmt.Println("==reverHead==:", reverHead)
	return dummyNode.Next
}
func reverseLinkedList(head *ListNode, dummyNode *ListNode) *ListNode {
	if head == nil {
		return dummyNode
	}

	reverseHead := reverseLinkedList(head.Next, dummyNode)
	reverseHead.Next = head
	head.Next = nil

	return head
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

//21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 输入：l1 = [1,2,4], l2 = [1,3,4]
	// 输出：[1,1,2,3,4,4]

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	dummyNode := &ListNode{
		Val: -101,
	}
	head := dummyNode
	node1 := l1
	node2 := l2

	for node2 != nil && node1 != nil {
		if node1.Val >= node2.Val {
			head.Next = node2
			node2 = node2.Next
		} else {
			head.Next = node1
			node1 = node1.Next
		}
		head = head.Next
	}
	if node2 != nil {
		head.Next = node2
	}

	if node1 != nil {
		head.Next = node1
	}
	return dummyNode.Next
}

//86. 分隔链表
func partition(head *ListNode, x int) *ListNode {
	/**
	  输入：head = [1,4,3,2,5,2], x = 3
	  输出：[1,2,2,4,3,5]
	*/

	if head == nil {
		return nil
	}

	headDummy := &ListNode{
		Val: 0,
	}
	tailDummy := &ListNode{
		Val: 0,
	}

	headDummy.Next = head
	head = headDummy
	tail := tailDummy

	// 0 5 1 2 3 4
	for head.Next != nil {
		fmt.Println("===head===:", head)
		if head.Next.Val < x {
			head = head.Next
		} else {
			tail.Next = head.Next
			tail = tail.Next
			head.Next = head.Next.Next
			fmt.Println("==tail==", tail)
		}

	}

	tail.Next = nil
	head.Next = tailDummy.Next
	fmt.Println("==head==", head)
	return headDummy.Next

}

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表
func sortList(head *ListNode) *ListNode {
	// 思路：归并排序，找中点和合并操作
	// 输入：head = [4,2,1,3]
	// 输出：[1,2,3,4]
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	// 找到中点
	midNode := findMiddle(head)
	rightHead := midNode.Next
	midNode.Next = nil

	// fmt.Println("==midNode==", midNode)
	// fmt.Println("==midNode==",midNode)
	// 两边分别排序

	leftHead := mergeSort(head)
	rightHead = mergeSort(rightHead)

	// 合并
	fmt.Println("==leftHead==", head)
	fmt.Println("==rightHead==", rightHead)
	sortedHead := mergeLeftRight(leftHead, rightHead)
	fmt.Println("==sortedHead==", sortedHead)
	return sortedHead
}

//合并两个有序的链表
func mergeLeftRight(leftHead *ListNode, rightHead *ListNode) *ListNode {
	if leftHead == nil {
		return rightHead
	}
	if rightHead == nil {
		return leftHead
	}

	dummyNode := &ListNode{
		Val: 0,
	}
	head := dummyNode
	for leftHead != nil && rightHead != nil {
		if leftHead.Val < rightHead.Val {
			head.Next = leftHead
			leftHead = leftHead.Next
		} else {
			head.Next = rightHead
			rightHead = rightHead.Next
		}
		head = head.Next
	}

	if leftHead == nil {
		head.Next = rightHead
	}
	if rightHead == nil {
		head.Next = leftHead
	}
	return dummyNode.Next
}

func findMiddle(head *ListNode) *ListNode {
	// 两个游标， 快的是慢的两倍
	// 12345
	// slow 1 2 3
	// fast 2 4 nil

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//143. 重排链表
func reorderList(head *ListNode) *ListNode {
	/**
	输入: head = [1,2,3,4,5]
	输出: [1,5,2,4,3]
	*/
	// 找中点
	midNode := findMiddle(head)
	rightHead := midNode.Next
	midNode.Next = nil
	// 反转后半部分
	revRightHead := reverseList3(rightHead)
	// 前后两部分合并
	orderHead := head
	for revRightHead != nil && head != nil {
		headNext := head.Next
		head.Next = revRightHead
		rightNext := revRightHead.Next
		revRightHead.Next = headNext
		head = headNext
		revRightHead = rightNext
	}
	return orderHead
}

func reverseList3(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//141. 环形链表
/**
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。
*/
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}
	var slow *ListNode
	var fast *ListNode
	slow = head
	fast = head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	// 1 2 3 4 5
	//此题 fast和slow从统一起点开始
	if head == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

//234. 回文链表
func isPalindrome(head *ListNode) bool {
	//快慢指针
	if head == nil {
		return false
	}
	//找到中点
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 如果 fast := head.Next，中点就是slow
	rightHead := slow.Next
	slow.Next = nil

	//反转后半段
	var pre *ListNode
	for rightHead != nil {
		next := rightHead.Next
		rightHead.Next = pre
		pre = rightHead
		rightHead = next
	}

	//反转后的和正向的对比
	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//138. 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	// 思路：1、hash 表存储指针，2、复制节点跟在原节点后面
	// 采用复制节点的方式
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		fmt.Println("==cur==", cur)
		copyNode := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = copyNode
		cur = copyNode.Next

	}

	//遍历 填入Random
	cur = head
	for cur != nil && cur.Next != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	//分开两个节点
	cur = head
	copyHead := head.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	return copyHead

}

func main() {

	//  [1,2,3,4,5]
	// 1 2
	// 4 3
	// 1423
	// data := []int{1, 2, 1, 2, 1}
	// head := buildLinkedList(data)

	node3 := &Node{
		Val: 3,
	}

	node2 := &Node{
		Val:    2,
		Next:   node3,
		Random: node3,
	}
	node1 := &Node{
		Val:    1,
		Next:   node2,
		Random: node3,
	}
	head := node1

	for head != nil {
		fmt.Println("1===node===:", head)
		head = head.Next
	}

	head = copyRandomList(node1)
	fmt.Println("===head===:", head)
	for head != nil {
		fmt.Println("2===node===:", head)
		head = head.Next
	}

}
