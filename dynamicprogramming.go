package algorithm

import "math"

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
