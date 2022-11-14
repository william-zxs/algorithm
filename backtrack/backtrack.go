package backtrack

import (
	"sort"
	"strings"
)

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

//77. 组合
func combine(n int, k int) [][]int {
	//子集问题的减枝
	if n == 1 && k == 1 {
		return [][]int{{1}}
	}

	result := make([][]int, 0)
	track := make([]int, 0)
	var trackback func(track []int, start int)
	trackback = func(track []int, start int) {
		//加入结果集
		if len(track) == k {
			result = append(result, append([]int{}, track...))
		}
		//start 从1开始
		for i := start; i <= n; i++ {
			//做选择
			track = append(track, i)
			//递归
			trackback(track, i+1)
			//撤销选择
			track = track[:len(track)-1]
		}
	}
	//start 从1开始
	trackback(track, 1)
	return result
}

//90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	track := make([]int, 0)
	// used := make([]bool,len(nums))
	result := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))
	var backtrack func(start int)
	backtrack = func(start int) {
		result = append(result, append([]int{}, track...))

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return result
}

//40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	sort.Sort(sort.IntSlice(candidates))
	var backtrack func(start int, sum int)

	backtrack = func(start int, sum int) {
		//先判断是否满足条件
		if sum == target {
			result = append(result, append([]int{}, track...))
		} else if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			//剪枝
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			track = append(track, candidates[i])
			sum += candidates[i]

			backtrack(i+1, sum)
			sum -= candidates[i]
			track = track[:len(track)-1]

		}

	}
	backtrack(0, 0)
	return result
}

//47. 全排列 II
func permuteUnique(nums []int) [][]int {
	//排序
	//剪枝
	//收集树的最后一层节点

	track := make([]int, 0)
	useMap := make(map[int]int, 0)
	result := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))

	for _, n := range nums {
		useMap[n] = useMap[n] + 1
	}

	var trackback func()
	trackback = func() {
		//判断是否满足要求
		if len(track) == len(nums) {
			result = append(result, append([]int{}, track...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if useMap[nums[i]] == 0 {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			//choice
			track = append(track, nums[i])
			useMap[nums[i]] = useMap[nums[i]] - 1
			trackback()
			track = track[:len(track)-1]
			useMap[nums[i]] = useMap[nums[i]] + 1

		}

	}
	trackback()
	return result
}

//39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	track := make([]int, 0)
	result := make([][]int, 0)
	var trackback func(start int, sum int)
	trackback = func(start int, sum int) {
		if sum == target {
			result = append(result, append([]int{}, track...))
			return
		} else if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			//选择
			track = append(track, candidates[i])
			sum += candidates[i]
			trackback(i, sum)
			sum -= candidates[i]
			track = track[:len(track)-1]
		}

	}
	trackback(0, 0)
	return result
}

//46. 全排列
func permute_(nums []int) [][]int {

	res := make([][]int, 0)
	var backtrack func(track []int, used []bool)

	backtrack = func(track []int, used []bool) {

		if len(track) == len(nums) {
			res = append(res, append(make([]int, 0), track...))
			return
		}

		for i := 0; i < len(nums); i++ {
			//做选择
			if used[i] == true {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			//  递归
			backtrack(track, used)
			//撤销选择
			track = track[:len(track)-1]
			used[i] = false

		}
	}
	track := make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(track, used)
	return res
}

//46. 全排列

func permute_3(nums []int) [][]int {
	//遍历决策树，暴力遍历，时间复杂度是O(n!)
	//但是必须说明的是，不管怎么优化，都符合回溯框架，
	//而且时间复杂度都不可能低于 O(N!)，
	//因为穷举整棵决策树是无法避免的。
	//这也是回溯算法的一个特点，
	//不像动态规划存在重叠子问题可以优化，
	//回溯算法就是纯暴力穷举，复杂度一般都很高。
	track := make([]int, 0)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			res = append(res, append(make([]int, 0), track...))
			return
		}
		for i := 0; i < len(nums); i++ {
			//做选择
			if used[i] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			//递归
			backtrack()
			//撤销选择
			used[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return res
}

//51. N 皇后
func solveNQueens_(n int) [][]string {
	track := make([]string, 0)
	result := make([][]string, 0)
	posSlice := make([]byte, n)
	for i := 0; i < len(posSlice); i++ {
		posSlice[i] = '.'
	}
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			//结束条件
			if len(track) == n {
				result = append(result, append(make([]string, 0), track...))
			}
			return
		}
		for j := 0; j < n; j++ {
			if !isOk(i, j, n, track) {
				continue
			}

			//做选择
			posSlice[j] = 'Q'
			line := string(posSlice)
			posSlice[j] = '.'
			track = append(track, line)
			//递归
			backtrack(i + 1)
			//撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return result
}

func isOk(i, j, n int, track []string) bool {
	//横 不用判断
	//上
	for l := i - 1; l >= 0; l-- {
		if track[l][j] == 'Q' {
			return false
		}
	}
	//右上
	for r, t := j+1, i-1; r < n && t >= 0; r, t = r+1, t-1 {
		if track[t][r] == 'Q' {
			return false
		}
	}
	//左上
	for l, t := j-1, i-1; l >= 0 && t >= 0; l, t = l-1, t-1 {
		if track[t][l] == 'Q' {
			return false
		}
	}
	return true
}
