package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

//24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	nextHead := next.Next
	next.Next = head
	head.Next = swapPairs(nextHead)
	return next
}
