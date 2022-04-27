package algorithm

//1、当输入规模n比较小的时候，应该使用选择排序或者时插入排序（插入排序通常会比选择排序少一些比较的次数，但是选择排序会比插入排序少挪动的次数）；
//2、当输入的数据基本有序的时候，应该使用插入排序或者是随机的快速排序。
//3、当n较大，要求速度较快的时候可以使用nlgn的算法比如：堆排序、快速排序和归并排序。堆排序的空间复杂度为lO(1)、而快速排序的空间复杂度是O(logn)(递归使用栈的情况)，而归并排序的时间复杂度最大是O(n).

//快排
//时间复杂度最好nlogn, 最差n2
func QuickSort(nums []int) {
	// 当数据越偏向于有序的时候，时间复杂度越接近n^2
	quickSort(nums, 0, len(nums)-1)
}
func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	endValue := nums[end]
	j := start
	for i := start; i < end; i++ {
		if nums[i] < endValue {
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

//归并排序
//时间复杂度是nlogn
//如何用公式计算时间复杂度
func MergeSort(nums []int) []int {
	//不断二分，最后比较两值，合并成一个有序数组，返回，然后再合并left和right的返回

	if len(nums) <= 1 {
		return nums
	}
	m := len(nums) / 2
	l := MergeSort(nums[:m])
	r := MergeSort(nums[m:])
	return merge(l, r)
}

func merge(left, right []int) []int {
	res := []int{}
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}

//堆排序
func HeapSort(a []int) []int {
	//构建一个大根堆
	for i := len(a)/2 - 1; i >= 0; i-- {
		sink(a, i, len(a))
	}

	//交换根节点和最后一个值
	for i := len(a) - 1; i >= 1; i-- {
		swap(a, 0, i)
		sink(a, 0, i)
	}
	return a
}
func sink(a []int, i int, length int) {
	for {
		l := i*2 + 1
		r := i*2 + 2
		idx := i
		if l < length && a[l] > a[idx] {
			idx = l
		}
		if r < length && a[r] > a[idx] {
			idx = r
		}

		if idx == i {
			break
		}
		swap(a, i, idx)
		i = idx
	}
}

func BubblingSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
	return nums
}
