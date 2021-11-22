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

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	p := root
	for p != nil {
		if val < p.Val {

			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				return root
			}
			p = p.Left
		} else {

			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				return root
			}
			p = p.Right
		}
	}
	return root
}

// 450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			//左右都有子节点的情况，把左节点移动到右节点的最后一个左子节点下
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}

//110. 平衡二叉树
/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/
func isBalanced(root *TreeNode) bool {
	//递归 计算深度
	if root == nil {
		return true
	}
	res := isBalancedHelper(root)
	if res == -1 {
		return false
	} else {
		return true
	}
}

func isBalancedHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := isBalancedHelper(root.Left)
	right := isBalancedHelper(root.Right)
	if right-left > 1 || left-right > 1 || left == -1 || right == -1 {
		return -1
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
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

	res := isValidBST(node1)
	fmt.Println("==res==:", res)
}
