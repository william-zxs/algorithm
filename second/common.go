package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
