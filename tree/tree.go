package tree

import "math"

//树问题的解决，可以用递归就用递归，会降低问题的复杂性

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	//dfs
	_, _, ok := helper(root)
	return ok
}

func helper(root *TreeNode) (min, max int, ok bool) {

	if root == nil {
		min = math.MaxInt
		max = math.MinInt
		ok = true
		return
	}

	leftMin, leftMax, leftOk := helper(root.Left)
	rightMin, rightMax, rightOk := helper(root.Right)

	if root.Left == nil {
		min = root.Val

	}

	if !leftOk || !rightOk {
		ok = false
		return
	}

	if root.Val > leftMax && root.Val < rightMin {
		min = leftMin
		max = rightMax
		if root.Left == nil {
			min = root.Val
		}

		if root.Right == nil {
			max = root.Val
		}

		ok = true
		return
	} else {
		ok = false
	}

	return

}

//98. 验证二叉搜索树
func isValidBST2(root *TreeNode) bool {
	//递归的方式
	return helper2(root, math.MinInt, math.MaxInt)

}

func helper2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return helper2(root.Left, min, root.Val) && helper2(root.Right, root.Val, max)
}

//98. 验证二叉搜索树
func isValidBST3(root *TreeNode) bool {
	//中序遍历
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

//701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	dummy := root

	for root != nil {
		if root.Val < val {
			if root.Right == nil {
				root.Right = &TreeNode{Val: val}
				return dummy
			}
			root = root.Right
		} else {
			if root.Left == nil {
				root.Left = &TreeNode{Val: val}
				return dummy
			}
			root = root.Left
		}
	}
	return dummy
}

//450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	right := root.Right
	for right.Left != nil {
		right = right.Left
	}
	right.Left = root.Left
	return root.Right
}

//110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	_, res := isBalancedHelper(root)
	return res
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, resl := isBalancedHelper(root.Left)
	r, resr := isBalancedHelper(root.Right)
	if !resl || !resr {
		return 0, false
	}
	if l-r > 1 || r-l > 1 {
		return 0, false
	}

	if l > r {
		return l + 1, true
	}
	return r + 1, true
}

//104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	//通过遍历树的思想解决问题
	var res int
	var depth int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		//前序位置
		depth++
		if root.Left == nil && root.Right == nil {
			if depth > res {
				res = depth
			}
		}
		//递归左右子树
		traverse(root.Left)
		traverse(root.Right)
		//后序位置
		depth--
	}
	traverse(root)
	return res
}

//543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxDiameter int

	//遍历二叉树，利用返回值，在后序位置写逻辑
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := maxDepth(root.Left)
		r := maxDepth(root.Right)
		//后序位置
		//计算直径
		if l+r > maxDiameter {
			maxDiameter = l + r
		}
		//返回最大深度
		if l > r {
			return l + 1
		}
		return r + 1
	}
	maxDepth(root)
	return maxDiameter
}
