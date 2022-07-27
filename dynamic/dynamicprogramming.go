package dynamic

import (
	"math"
)

//Matrix DP (10%)
//Sequence (40%)
//Two Sequences DP (40%)
//Backpack (10%)

//Matrix DP (10%)
//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	//动态规划，找到状态转移方程
	//f[i][j] = min(f[i-1][j],f[i-1][j-1]) + triangle[i][j]
	//同时处理一下边界值就可以了
	//还可以优化一下内存占用
	n := len(triangle)

	//构建一个结果集
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	//遍历所有的节点，求最小值
	//初始值
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				f[i][0] = f[i-1][0] + triangle[i][j]
				continue
			}
			if j == i {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
				continue
			}
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
	}
	minV := math.MaxInt
	//遍历最后一列
	for i := 0; i < n; i++ {
		minV = min(minV, f[n-1][i])
	}
	return minV
}

func min(n, m int) int {
	if m < n {
		return m
	}
	return n
}

//64. 最小路径和 中等
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					f[0][0] = grid[0][0]
				} else {
					f[0][j] = f[0][j-1] + grid[i][j]
				}
				continue
			}
			if j == 0 {
				f[i][0] = f[i-1][0] + grid[i][0]
				continue
			}
			f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}
	return f[m-1][n-1]

}

//62. 不同路径 中等
func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[0][0] = 1
				continue
			}
			if i == 0 {
				f[0][j] = f[0][j-1]
				continue
			}
			if j == 0 {
				f[i][0] = f[i-1][0]
				continue
			}
			f[i][j] = f[i][j-1] + f[i-1][j]
		}
	}
	return f[m-1][n-1]
}

//63. 不同路径 II 中等
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			if i == 0 && j == 0 {
				obstacleGrid[0][0] = 1
				continue
			}

			if i == 0 {
				obstacleGrid[0][j] = obstacleGrid[0][j-1]
				continue
			}
			if j == 0 {
				obstacleGrid[i][0] = obstacleGrid[i-1][0]
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
		}
	}
	return obstacleGrid[m-1][n-1]
}

//Sequence (40%)
//70. 爬楼梯 简单
func climbStairs(n int) int {
	a := 1
	b := 1
	c := 0
	for i := 2; i < n+1; i++ {
		c = a + b
		a = b
		b = c
	}

	return b
}

//55. 跳跃游戏 中等
func canJump(nums []int) bool {
	// 也可以用贪心算法
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if f[j] == false {
				continue
			}
			if nums[j] >= i-j {
				f[i] = true
				break
			}
		}
	}

	return f[len(nums)-1]
}

//45. 跳跃游戏 II 中等
func jump(nums []int) int {
	//使用贪心算法 时间复杂度可以降低到n todo
	f := make([]int, len(nums))
	for i := 0; i < len(f); i++ {
		f[i] = math.MaxInt
	}
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(nums)-1]
}

//131. 分割回文串 中等
func Partition(s string) (res [][]string) {
	//动态规划+回溯
	l := len(s)
	f := make([][]bool, l)
	for i := 0; i < l; i++ {
		f[i] = make([]bool, l)
		for j := 0; j < l; j++ {
			f[i][j] = true
		}
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var dfs func(i int)
	split := []string{}
	dfs = func(i int) {
		if i == l {
			res = append(res, append([]string{}, split...))
			return
		}

		for j := i; j < l; j++ {
			if f[i][j] {
				split = append(split, s[i:j+1])
				dfs(j + 1)
				split = split[:len(split)-1]
			}
		}
	}
	dfs(0)
	return
}

//300. 最长递增子序列 中等
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)
	dp = append(dp, 1)
	maxL := 1
	for i := 1; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxL = max(maxL, dp[i])
	}
	return maxL
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

//139. 单词拆分 中等
func wordBreak(s string, wordDict []string) bool {
	existMap := make(map[string]bool, 0)
	for _, word := range wordDict {
		existMap[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && existMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//322. 零钱兑换
//中等
func coinChange(coins []int, amount int) int {
	// 需要再优化
	dp := make(map[int]int, 0)
	dp[0] = 0

	var min func(amount int) (minCount int)
	min = func(amount int) (minCount int) {
		minCount = -1
		for _, v := range coins {
			if amount-v < 0 {
				continue
			} else if amount-v == 0 {
				minCount = 0
				return
			}
			if dp[amount-v] > 0 {
				if minCount == -1 || minCount > dp[amount-v] {
					minCount = dp[amount-v]
				}
			} else {
				continue
			}
		}
		return
	}

	for i := 1; i <= amount; i++ {
		minCount := min(i)
		if minCount == -1 {
			dp[i] = -1
		} else {
			dp[i] = minCount + 1
		}
	}
	return dp[amount]
}
