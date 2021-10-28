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

func main() {
	// triangle := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// res := minPathSum(triangle)
	// res := uniquePaths(3, 7)
	res := jump([]int{3, 4, 3, 2, 5, 4, 3})
	fmt.Println("==res==:", res)
}
