package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return dummy.Next
}

//86. 分隔链表
func partition(head *ListNode, x int) *ListNode {
	lDummy := &ListNode{}
	rDummy := &ListNode{}
	l := lDummy
	r := rDummy
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			r.Next = head
			r = r.Next
		}
		head = head.Next
	}
	l.Next = rDummy.Next
	r.Next = nil
	return lDummy.Next
}

func main() {

}
