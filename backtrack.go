package main

import "fmt"

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

func main() {

	data := []int{1, 2, 3}
	res := subsets(data)
	fmt.Println("==res==", res)
}
