package main

import (
	"fmt"
	"math"
)

/*
二叉搜索树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 98. 验证二叉搜索树
func isValidBST2(root *TreeNode) bool {
	/*
		左节点小于父节点，右节点大于父节点
	*/

	flag := true

	var doWork func(root *TreeNode) (int, int)
	doWork = func(root *TreeNode) (int, int) {

		if flag == false {
			return 0, 0
		}
		var left, right int
		if root == nil {
			left, right = math.MaxInt32, math.MinInt32
			return left, right
		}
		if root.Left == nil && root.Right == nil {
			return root.Val, root.Val
		}

		// if root.Val <= root.Right {

		// }

		leftMin, leftMax := doWork(root.Left)
		rightMin, rightMax := doWork(root.Right)

		if leftMax >= root.Val {
			flag = false
		}
		if rightMin <= root.Val {
			flag = false
		}

		if leftMin > leftMax {
			leftMin = root.Val
		}

		if rightMax < rightMin {
			rightMax = root.Val
		}
		fmt.Println("==leftMin==", leftMin, " ==rightMax==", rightMax, " ==root.Val==", root.Val)
		return leftMin, rightMax
	}

	doWork(root)
	return flag
}

//  98. 验证二叉搜索树  中序遍历
func isValidBST(root *TreeNode) bool {
	// 遍历左子树 保存到 切片中

	stack := []*TreeNode{}
	inorder := math.MinInt64

	for len(stack) > 0 || root != nil {
		fmt.Println("==stack==:", stack)
		fmt.Println("==root==:", root)
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		// fmt.Println("==root==:", root)
		// fmt.Println("==stack==:", stack)
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true

}

// 98. 验证二叉搜索树
func isValidBST3(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func main() {

	// [5,1,4,null,null,3,6]
	/*
			5
		  1   4
		     3 6
	*/

	node7 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	node6 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node3 := &TreeNode{
		Val:   4,
		Left:  node6,
		Right: node7,
	}
	node2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	node1 := &TreeNode{
		Val:   5,
		Left:  node2,
		Right: node3,
	}

	// [2,1,3]
	// node3 := &TreeNode{
	// 	Val:   3,
	// 	Left:  nil,
	// 	Right: nil,
	// }

	// node2 := &TreeNode{
	// 	Val:   1,
	// 	Left:  nil,
	// 	Right: nil,
	// }

	// node1 := &TreeNode{
	// 	Val:   2,
	// 	Left:  node2,
	// 	Right: node3,
	// }

	// node1 := &TreeNode{
	// 	Val:   2147483647,
	// 	Left:  nil,
	// 	Right: nil,
	// }

	// res := isValidBST(node1)
	res := node1
	fmt.Println("==res1==:", res)
	res := node2
	fmt.Println("==res2==:", res)
}
