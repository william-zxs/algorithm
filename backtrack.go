package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
回溯算法

参考：
	https://zhuanlan.zhihu.com/p/93530380
	labuladong

	https://greyireland.gitbook.io/algorithm-pattern/suan-fa-si-wei/backtrack
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

func combinationSum(candidates []int, target int) [][]int {
	list := make([]int, 0)
	result := make([][]int, 0)
	sort.Ints(candidates)
	combinationSumHelper(candidates, list, &result, target)
	return result
}

//39. 组合总和
func combinationSumHelper(candidates []int, list []int, result *[][]int, target int) {
	sumVal := sumList(list)
	if sumVal == target {
		data := make([]int, len(list))
		copy(data, list)
		*result = append(*result, data)
		return
	} else if sumVal > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		if len(list) > 0 && list[len(list)-1] > candidates[i] {
			continue
		}
		list = append(list, candidates[i])
		combinationSumHelper(candidates, list, result, target)
		list = list[:len(list)-1]
	}

}

func sumList(list []int) int {
	sumVal := 0
	for i := 0; i < len(list); i++ {
		sumVal += list[i]
	}
	return sumVal
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	path := make([]byte, 0)
	if len(digits) == 0 {
		return result
	}

	dataMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	letterCombinationsHelper(&result, path, digits, 0, dataMap)
	return result
}

func letterCombinationsHelper(result *[]string, path []byte, digits string, pos int, dataMap map[string]string) {
	if len(path) == len(digits) {
		data := make([]byte, len(path))
		copy(data, path)
		*result = append(*result, string(data))
		return
	}

	numStr := string(digits[pos])
	if str, ok := dataMap[numStr]; ok {
		for y := 0; y < len(str); y++ {
			path = append(path, str[y])
			letterCombinationsHelper(result, path, digits, pos+1, dataMap)
			path = path[:len(path)-1]
		}
	}
}

// 131. 分割回文串
func partition(s string) [][]string {
	l := len(s)
	pd := make([][]bool, l)

	result := make([][]string, 0)

	//预处理
	for i := 0; i < l; i++ {
		inpd := make([]bool, l)
		for y := 0; y < l; y++ {
			inpd[y] = true
		}
		pd[i] = inpd
	}

	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			pd[i][j] = s[i] == s[j] && pd[i+1][j-1]
		}
	}

	// 路径
	list := make([]string, 0)
	//回溯 递归
	var partitionHelper func(i int)
	partitionHelper = func(i int) {
		if i == l {
			num := len(list)
			list2 := make([]string, num)
			copy(list2, list)
			result = append(result, list2)
			return
		}

		for j := i; j < l; j++ {
			if pd[i][j] {
				list = append(list, s[i:j+1])
				partitionHelper(j + 1)
				list = list[:len(list)-1]
			}
		}
	}
	partitionHelper(0)
	return result
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	list := make([]string, 0)
	restoreIpAddressesHelper(s, 0, list, &result)
	return result
}

func restoreIpAddressesHelper(s string, pos int, list []string, result *[]string) {

	if len(list) == 4 {
		if pos == len(s) {
			for i := 0; i < 4; i++ {
				if !checkLegal(list[i]) {
					return
				}
			}
			*result = append(*result, strings.Join(list, "."))
		}
		return
	}
	for i := pos; i < len(s); i++ {
		list = append(list, s[pos:i+1])
		restoreIpAddressesHelper(s, i+1, list, result)
		list = list[:len(list)-1]
	}
}

func checkLegal(subS string) bool {
	if len(subS) == 0 {
		return false
	}
	if len(subS) > 1 && string(subS[0]) == "0" {
		return false
	}

	num, error := strconv.Atoi(subS)
	if error == nil {
		if num >= 0 && num <= 255 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

func main() {

	data := "aab"
	res := partition(data)
	fmt.Println("==res==", res)
}
