package iteration

type ListNode struct {
	Val  int
	Next *ListNode
}

//24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}

	var pre *ListNode
	pre = dummy
	for head != nil && head.Next != nil {

		nnext := head.Next.Next
		head.Next.Next = head
		pre.Next = head.Next
		head.Next = nnext

		pre = head
		head = nnext
	}
	return dummy.Next
}
