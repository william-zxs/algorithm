package main

import (
	"fmt"
)

//排序
/*
排序分类：
快排
归并
堆排
*/

//快排
func QuickSort(nums []int) []int {
	//最重要的核心就是分治
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	// 快排的思路是用 pivot 区分大于和小于的两部分，分治法，不断缩小需要处理的数据，直到数据被分成了1、2、3个
	if start >= end {
		return
	}

	p := nums[end]
	j := start
	for i := 0; i < len(nums); i++ {
		if nums[i] < p {
			//交换 i 和 j
			swap(nums, i, j)
			j++
		}
	}
	swap(nums, j, end)
	// return
	//分治
	quickSort(nums, start, j-1) // [2,1,3],0,1
	quickSort(nums, j+1, end)   // [2,1,3],3,2
	return
}

// 一个点，切片是一个结构体，作为参数也是值传递，值记录了数组的地址，但是当删除和增加的时候，需要传递指针
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//归并排序
func MergeSort(nums []int) []int {
	//归并也是分治，分成多个小段
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func main() {
	data := []int{2, 1, 3, 6, 4, 10, 7}

	res := QuickSort(data)
	fmt.Println("==res==:", res)

}
