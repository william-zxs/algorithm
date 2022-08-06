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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//95. 不同的二叉搜索树 II
//中等
func GenerateTrees(n int) []*TreeNode {
	//迭代的思想
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftList := helper(start, i-1)
		rightList := helper(i+1, end)

		for _, left := range leftList {
			for _, right := range rightList {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}

	}
	return res
}
