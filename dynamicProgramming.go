package main

import (
	"fmt"
	"math"
)

//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	// 在triangle中修改

	triangle[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		//每行第一个的minimumTotal 等于上一行第一个加triangle[i][0]
		triangle[i][0] = triangle[i-1][0] + triangle[i][0]

		//中间数据 [1,i-1]
		for n := 1; n < i; n++ {
			triangle[i][n] = min(triangle[i-1][n], triangle[i-1][n-1]) + triangle[i][n]
		}

		//每行最后一个 等于 上一行最后一个加本行最后一个的值
		triangle[i][i] = triangle[i-1][i-1] + triangle[i][i]

	}

	cur := math.MaxInt32
	for i := 0; i < len(triangle); i++ {
		cur = min(triangle[len(triangle)-1][i], cur)
	}
	fmt.Println("==triangle==", triangle)
	return cur
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

//64. 最小路径和
func minPathSum(grid [][]int) int {

	cache := make([][]int, len(grid))
	for i := range cache {
		cache[i] = make([]int, len(grid[0]))
	}
	var minxy func(x, y int) int
	minxy = func(x, y int) int {

		if cache[x][y] != 0 {
			return cache[x][y]
		}
		if x == len(grid)-1 && y == len(grid[0])-1 {
			return grid[x][y]
		}
		if x == len(grid)-1 {
			return minxy(x, y+1) + grid[x][y]
		}
		if y == len(grid[0])-1 {
			return minxy(x+1, y) + grid[x][y]
		}

		val := min(minxy(x+1, y), minxy(x, y+1)) + grid[x][y]
		cache[x][y] = val
		return val
	}

	return minxy(0, 0)
}

// 62. 不同路径
// 也可以用数学的方式，直接计算排列组合
func uniquePaths(m int, n int) int {
	// [[1,2],[1,2],[1,2]]  m =3 n =2
	// 重点是找到状态转移方程
	// f(x, y) = f(x, y-1) + f(x-1, y)
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if x == 0 && y == 0 {
				path[0][0] = 1
			} else if x == 0 {
				path[0][y] = path[0][y-1]
			} else if y == 0 {
				path[x][0] = path[x-1][0]
			} else {
				path[x][y] = path[x][y-1] + path[x-1][y]
			}

		}
	}
	return path[m-1][n-1]
}

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if obstacleGrid[x][y] == 1 {
				path[x][y] = 0
				continue
			}
			if x == 0 && y == 0 {
				path[0][0] = 1
			} else if x == 0 {
				path[0][y] = path[0][y-1]
			} else if y == 0 {
				path[x][0] = path[x-1][0]
			} else {
				path[x][y] = path[x][y-1] + path[x-1][y]
			}
		}
	}
	return path[m-1][n-1]
}

//70. 爬楼梯
func climbStairs(n int) int {
	/*
			假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

		每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

		注意：给定 n 是一个正整数。
	*/

	//动态规划，要把大问题转化成小问题，先找到转移方程
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

//55. 跳跃游戏
func canJump(nums []int) bool {
	zeroI := -1
	for i := len(nums) - 2; i >= 0; i-- {

		if nums[i] == 0 && zeroI == -1 {
			zeroI = i
			continue
		}
		if zeroI != -1 {
			if nums[i]-nums[zeroI] > zeroI-i {
				zeroI = -1
			}
		}
	}
	if zeroI != -1 {
		return false
	}
	return true
}

//55. 跳跃游戏
// 动态规划的解决办法
func canJump2(nums []int) bool {
	// 思路：看最后一跳
	// 状态：f[i] 表示是否能从0跳到i
	// 推导：f[i] = OR(f[j],j<i&&j能跳到i) 判断之前所有的点最后一跳是否能跳到当前点
	// 初始化：f[0] = 0
	// 结果： f[n-1]
	if len(nums) == 0 {
		return true
	}
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] == true && nums[j]+j >= i {
				f[i] = true
			}
		}
	}
	return f[len(nums)-1]
}

// 45. 跳跃游戏 II
func jump(nums []int) int {
	//维护一个可跳最远的长度，直到超过最后一个
	if len(nums) <= 1 {
		return 0
	}
	maxI := 0     //3 5
	count := 0    //1
	stepMaxI := 0 //3
	for i := 0; i < len(nums); i++ {
		maxI = max(nums[i]+i, maxI)
		if maxI >= len(nums)-1 {
			count++
			break
		}
		if i >= stepMaxI {
			count++
			stepMaxI = maxI
		}
	}
	return count
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 132. 分割回文串 II
/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
返回符合要求的 最少分割次数 。
*/
func minCut(s string) int {
	// state: f[i] "前i"个字符组成的子字符串需要最少几次cut(个数-1为索引)
	// function: f[i] = MIN{f[j]+1}, j < i && [j+1 ~ i]这一段是一个回文串
	// intialize: f[i] = i - 1 (f[0] = -1)
	// answer: f[s.length()]
	// aabb
	n := len(s) // 4
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
		for j := range g[i] {
			g[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		// i= 3
		// i = 2
		for j := i + 1; j < n; j++ {
			// j = 4
			// j = 3
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
			//g[2][3] = s[2] == s[3] &&g[3][2]

		}
	}

	f := make([]int, n)
	for i := range f {
		if g[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if g[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	return f[n-1]
}

// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {
	//思路 动态规划，先找到转移方程
	// f[i] 到i为止的最长子序列长度
	// f[i] = max(f[j])+1  a[j]<a[i]

	// e.g.  [1,3,2,5,4]
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	maxL := 0
	for i := 0; i < len(nums); i++ {
		maxL = max(f[i], maxL)
	}
	return maxL
}

//131. 分割回文串
/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
*/
func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	fmt.Println("==f==:", f)
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				fmt.Println("==ans==0:", ans)
				fmt.Println("==splits==1:", splits)
				splits = splits[:len(splits)-1]
				fmt.Println("==splits==2:", splits)
			}
		}
	}
	dfs(0)
	return

}

// 5. 最长回文子串
func longestPalindrome(s string) string {
	// 动态规划 先初始化所有的状态
	// f(x,y) = f(x+1,y-1) && s[x]==s[y]
	n := len(s)
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for y := i + 1; y < n; y++ {
			f[i][y] = f[i+1][y-1] && s[i] == s[y]
		}
	}
	maxS := ""
	for i := 0; i < n; i++ {
		for y := 0; y < n; y++ {
			if y >= i && f[i][y] {
				if y-i+1 > len(maxS) {
					maxS = s[i : y+1]
				}
			}
		}
	}
	return maxS
}

func main() {
	// res := minCut("aabb")
	// partition("aab")

	// splits := []string{"a", "b", "c", "d"}
	// splits = splits[:len(splits)-1]

	res := partition("aab")
	fmt.Println("==res==:", res)
}
