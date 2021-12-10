package main

import "sort"

// 28. 实现 strStr()
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

//78. 子集
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	list := make([]int, 0)
	subsetsHelper(0, nums, list, &result)
	return result
}

func subsetsHelper(pos int, nums []int, list []int, result *[][]int) {

	data := make([]int, len(list))
	copy(data, list)
	*result = append(*result, data)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		subsetsHelper(i+1, nums, list, result)
		list = list[:len(list)-1]
	}
}
