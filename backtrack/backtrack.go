package backtrack

import "strings"

//回溯法

//46. 全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(nums []int, track []int, used []bool)
	backtrack = func(nums []int, track []int, used []bool) {
		if len(track) == len(nums) {
			result = append(result, append([]int{}, track...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtrack(nums, track, used)
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backtrack(nums, track, used)

	return result
}

func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	result := make([][]string, 0)

	var backtrack func(m int)
	backtrack = func(m int) {
		if m == n {
			//保存结果
			res := []string{}
			for i := 0; i < len(board); i++ {
				res = append(res, strings.Join(board[i], ""))
			}
			result = append(result, res)
			return
		}
		for i := 0; i < n; i++ {
			//校验
			if !IsValid(board, m, i) {
				continue
			}
			//选择
			board[m][i] = "Q"
			//递归
			backtrack(m + 1)
			//撤销
			board[m][i] = "."
		}
	}

	backtrack(0)
	return result
}

//51. N 皇后
func IsValid(board [][]string, m, i int) bool {
	//校验左上角
	for m, i := m-1, i-1; m >= 0 && i >= 0; m, i = m-1, i-1 {

		if board[m][i] == "Q" {
			return false
		}

	}
	//校验上面
	for m := m - 1; m >= 0; m = m - 1 {
		if board[m][i] == "Q" {
			return false
		}
	}
	//校验右上角
	for m, i := m-1, i+1; m >= 0 && i < len(board); m, i = m-1, i+1 {
		if board[m][i] == "Q" {
			return false
		}
	}
	return true
}

//78. 子集
func subsets(nums []int) [][]int {

	result := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		//前置
		result = append(result, append([]int{}, track...))
		for i := start; i < len(nums); i++ {
			//选择
			track = append(track, nums[i])
			//递归
			backtrack(nums, i+1)
			//撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, 0)
	return result
}
