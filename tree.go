package main

import "fmt"

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	levelData := make([]int, 0)
	levelData = append(levelData, root.Val)

	left := levelOrder(root.Left)
	right := levelOrder(root.Right)

}

func doWork(root *TreeNode, container [][]int, depth int) [][]int {
	if root == nil {
		return nil
	}
	//depth 从 -1 开始
	depth++
	if len(container) >= depth+1 {
		container[depth] = append(container[depth], root.Val)
	}
	levelData := make([]int, 0)
	levelData = append(levelData, root.Val)

	left := doWork(root.Left, container)
	right := doWork(root.Right, container)
}

func main() {
	// 可以为nil的情况
	// num := -(1 << 31)
	// fmt.Println("num==", num)
	testMax()
}
