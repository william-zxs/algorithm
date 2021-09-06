package main

import (
	"fmt"
	"math"
)

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("===i===:%d ", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("===j===:%d ", j)
			if j == 7 {
				fmt.Println(" ")
				break
			}
		}
	}
}

func lengthDemo() {
	str1 := "william"
	fmt.Println(len(str1))
}

func strStr(haystack string, needle string) int {
	// 如果长度是0，返回0
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(haystack); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func swapInt(x, y *int) {
	fmt.Printf("==x==:%d\n", *x)
	m := *x + 1
	fmt.Printf("==x==:%d\n", m)

	*x, *y = *y, *x
}

func maxDepth(root *TreeNode) int {
	depth := getNodeDepth(root)
	return depth
}

func getNodeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftRes := getNodeDepth(root.Left)
	rightRes := getNodeDepth(root.Right)

	if leftRes > rightRes {
		return leftRes + 1
	}
	return rightRes + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res := getSonDepthGap(root)
	if res == -1 {
		return false
	} else {
		return true
	}
}

func getSonDepthGap(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftRes := getSonDepthGap(root.Left)
	rightRes := getSonDepthGap(root.Right)

	if leftRes == -1 || rightRes == -1 {
		return -1
	}

	if leftRes-rightRes > 1 || rightRes-leftRes > 1 {
		return -1
	}

	if leftRes > rightRes {
		return leftRes + 1
	}
	return rightRes + 1
}

type ResultType struct {
	SinglePath int
	MaxSum     int
}

func maxPathSum(root *TreeNode) int {
	data := maxPathSumMax(root)
	fmt.Printf("==SinglePath===%d  ==MaxSum===%d \n", data.SinglePath, data.MaxSum)
	return data.MaxSum
}

func maxPathSumMax(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxSum:     -(1 << 31),
		}
	}
	var resType ResultType

	left := maxPathSumMax(root.Left)
	right := maxPathSumMax(root.Right)

	biggerSinglePath := max(left.SinglePath, right.SinglePath)

	resType.SinglePath = max(biggerSinglePath+root.Val, 0)

	resType.MaxSum = max(max(left.MaxSum, right.MaxSum), left.SinglePath+right.SinglePath+root.Val)
	// 单边最大+max(root.Val,0)
	return resType
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testMax() {

	data3 := TreeNode{
		Val:   -2,
		Left:  nil,
		Right: nil,
	}

	data2 := TreeNode{
		Val:   -1,
		Left:  nil,
		Right: nil,
	}

	data1 := TreeNode{
		Val:   2,
		Left:  &data2,
		Right: &data3,
	}

	maxSum := maxPathSum(&data1)
	fmt.Println("==maxSum===:", maxSum)
}

// 计算公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := hasPoint(root, p, q)
	return result.commonNode

}

type ResultCommon struct {
	count      int
	commonNode *TreeNode
}

func hasPoint(root, p, q *TreeNode) ResultCommon {
	if root == nil {
		return ResultCommon{
			count:      0,
			commonNode: nil,
		}
	}

	left := hasPoint(root.Left, p, q)
	right := hasPoint(root.Right, p, q)

	if left.count == 2 {
		return left
	}
	if right.count == 2 {
		return right
	}
	count := 0
	if root.Val == p.Val || root.Val == q.Val {
		count += 1
	}
	count += left.count + right.count

	if count == 2 {
		return ResultCommon{
			count:      2,
			commonNode: root,
		}
	} else if count == 1 {
		return ResultCommon{
			count:      1,
			commonNode: nil,
		}
	} else {
		return ResultCommon{
			count:      0,
			commonNode: nil,
		}
	}

}

// 公共祖先 优化版
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

//层序遍历
// 主要的思路就是通过队列一层一层的去处理
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	container := make([][]int, 0)
	for len(queue) > 0 {
		length := len(queue)
		levelContainer := make([]int, 0)

		for i := 0; i < length; i++ {
			fmt.Printf("===length===%d====i===%d \n", length, i)
			// break
			node := queue[0]
			queue = queue[1:]
			levelContainer = append(levelContainer, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		// queue = queue[length:]
		container = append(container, levelContainer)

	}
	reverseCon := make([][]int, 0)
	for i := len(container) - 1; i >= 0; i-- {
		reverseCon = append(reverseCon, container[i])
	}
	return reverseCon
}

//锯齿形遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	container := make([][]int, 0)
	for len(queue) > 0 {
		length := len(queue)
		levelContainer := make([]int, 0)

		for i := 0; i < length; i++ {
			fmt.Printf("===length===%d====i===%d \n", length, i)
			// break
			node := queue[0]
			queue = queue[1:]
			levelContainer = append(levelContainer, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		// queue = queue[length:]
		container = append(container, levelContainer)

	}
	zigzagCon := make([][]int, 0)
	for i := 0; i < len(container); i++ {
		if i%2 == 0 {
			zigzagCon = append(zigzagCon, container[i])
		} else {
			list := make([]int, 0)
			for m := len(container[i]) - 1; m >= 0; m-- {
				list = append(list, container[i][m])
			}
			zigzagCon = append(zigzagCon, list)
		}
	}
	return zigzagCon
}

/**
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
**/

// 可以用分治法，也可以用中序遍历
// 验证二叉搜索树
type ValidBSTType struct {
	maxVal     int
	minVal     int
	isValidBST bool
}

func isValidBST(root *TreeNode) bool {
	res := doValidBST(root)
	return res.isValidBST
}

func doValidBST(root *TreeNode) ValidBSTType {
	if root == nil {
		return ValidBSTType{
			maxVal:     0,
			minVal:     0,
			isValidBST: true,
		}
	}

	left := doValidBST(root.Left)
	right := doValidBST(root.Right)
	//校验历史数据 如果不是BST 直接返回
	if !(left.isValidBST && right.isValidBST) {
		return ValidBSTType{
			maxVal:     0,
			minVal:     0,
			isValidBST: false,
		}
	}

	maxVal := right.maxVal
	minVal := left.minVal

	// 判断当前节点是否是BST
	if root.Left != nil {
		if root.Val <= left.maxVal {
			return ValidBSTType{
				maxVal:     0,
				minVal:     0,
				isValidBST: false,
			}
		}
	} else {
		minVal = root.Val
	}

	if root.Right != nil {
		if root.Val >= right.minVal {
			return ValidBSTType{
				maxVal:     0,
				minVal:     0,
				isValidBST: false,
			}
		}
	} else {
		maxVal = root.Val
	}

	return ValidBSTType{
		maxVal:     maxVal,
		minVal:     minVal,
		isValidBST: true,
	}
}

/**
验证 是否是二叉查找树  BST

中序遍历

**/
func isValidBST2(root *TreeNode) bool {
	// 处理的顺序是节点加入的顺序决定的
	// 中序遍历 找到循环的条件
	list := make([]*TreeNode, 0)
	// 判断条件
	smaller := math.MinInt64
	// 数据队列
	for len(list) > 0 || root != nil {
		// 中序遍历 从 最左节点开始
		for root != nil {
			list = append(list, root)
			root = root.Left
		}

		root = list[len(list)-1]
		list = list[:len(list)-1]

		if root.Val <= smaller {
			return false
		}

		smaller = root.Val
		root = root.Right
	}
	return true
}

// 二叉搜索树中的插入操作
// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//插入条件
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func main() {

	data2 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	data3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	data1 := TreeNode{
		Val:   2,
		Left:  &data2,
		Right: &data3,
	}
	res := isValidBST(&data1)
	fmt.Println(res)
}
