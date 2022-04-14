package algorithm

//快排
func QuickSort(nums []int) []int {
	//快排的思路是选取最后一个值为基准值r，遍历list和基准值比较，
	//一个记录最左边比基准值大的一个 j，当前遍历i比基准值要小就交换swap(j,i),
	//同时 j++,直到遍历完成，将基准值r和j互换，按照索引j的分成两份，递归
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	b := nums[end]
	j := start
	for i := start; i < end; i++ {
		if nums[i] < b {
			swap(nums, j, i)
			j++
		}
	}
	swap(nums, j, end)
	quickSort(nums, start, j-1)
	quickSort(nums, j+1, end)
}
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
