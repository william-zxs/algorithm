package main

import (
	"fmt"
	"sort"
)

/*
回溯算法
*/

// 78. 子集
func subsets(nums []int) [][]int {
	//决策树
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	val := make([]int, len(list))
	copy(val, list)
	*result = append(*result, val)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		fmt.Println("==list==", list)
		backtrack(nums, i+1, list, result)
		list = list[:len(list)-1]
	}
}

// 90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	sort.Ints(nums)
	backtrackWithDup(nums, 0, list, &result)
	return result
}

func backtrackWithDup(nums []int, pos int, list []int, result *[][]int) {
	//保存路径
	data := make([]int, len(list))
	copy(data, list)
	*result = append(*result, data)
	for i := pos; i < len(nums); i++ {
		//剪枝
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		//做选择
		list = append(list, nums[i])
		//递归
		backtrackWithDup(nums, i+1, list, result)
		//撤销选择
		list = list[:len(list)-1]
	}
}

// 46. 全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrackPermute(nums, list, &result)
	return result
}

func backtrackPermute(nums []int, list []int, result *[][]int) {
	if len(list) == len(nums) {
		data := make([]int, len(list))
		copy(data, list)
		*result = append(*result, data)
		return
	}
out:
	for i := 0; i < len(nums); i++ {
		for y := 0; y < len(list); y++ {
			if nums[i] == list[y] {
				continue out
			}
		}
		list = append(list, nums[i])
		backtrackPermute(nums, list, result)
		list = list[:len(list)-1]
	}
}

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	backtrackPermuteUnique(nums, list, &result, visited)
	return result
}
func backtrackPermuteUnique(nums []int, list []int, result *[][]int, visited []bool) {
	if len(list) == len(nums) {
		data := make([]int, len(list))
		copy(data, list)
		*result = append(*result, data)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrackPermuteUnique(nums, list, result, visited)
		visited[i] = false
		list = list[:len(list)-1]
	}
}

func main() {

	data := []int{1, 2, 3}
	res := permute(data)
	fmt.Println("==res==", res)
}
