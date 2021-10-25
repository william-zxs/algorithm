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
	for i := start; i < len(nums); i++ {
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

/*

一般都用数组来表示堆，i结点的父结点下标就为(i-1)/2。
它的左右子结点下标分别为2 * i + 1和2 * i + 2。如第0个结点左右子结点下标分别为1和2。

对于叶子节点，不用调整次序，根据满二叉树的性质，叶子节点比内部节点的个数多1.所以i=n/2 -1 ，不用从n开始。
*/
//堆排序
func HeapSort(nums []int) []int {
	//思路 先构建一个大根堆，交换根节点和最后一个节点，继续恢复大根堆

	//构建大根堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i, len(nums))
	}

	//排序
	// 最后一个和根节点置换
	for i := len(nums) - 1; i >= 1; i-- {
		//置换i 和 0
		swap(nums, 0, i)

		//恢复大根堆
		sink(nums, 0, i)
	}
	return nums
}

func sink(nums []int, i, len int) {
	for {
		l := 2*i + 1
		r := 2*i + 2
		idx := i

		if l < len && nums[l] > nums[i] {
			idx = l
		}

		if r < len && nums[r] > nums[idx] {
			idx = r
		}
		if idx == i {
			break
		}

		swap(nums, i, idx)

		i = idx
	}

}

func main() {
	data := []int{5, 1, 3, 2, 8, 3, 5, 6, 1, 3, 4, 5, 7, 7, 10, 3123, 4, 3, 535, 542, 12, 1, 293}

	res := HeapSort(data)
	fmt.Println("==res==:", res)

}
