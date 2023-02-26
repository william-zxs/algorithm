package main

//BFS 出现的常见场景好吧，问题的本质就是让你在一幅「图」中找到从起点 start 到终点 target 的最近距离，
//这个例子听起来很枯燥，但是 BFS 算法问题其实都是在干这个事儿

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
// 简单
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	list := make([]*TreeNode, 0)
	list = append(list, root)
	step := 1
	for len(list) > 0 {
		l := len(list)
		for i := 0; i < l; i++ {
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
			if list[i].Left == nil && list[i].Right == nil {
				return step
			}
		}
		step += 1
		list = list[l:]
	}
	return step
}

func main() {

}
