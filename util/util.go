package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func Max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
