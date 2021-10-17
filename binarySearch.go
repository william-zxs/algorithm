package main

import "fmt"

//704. 二分查找
func search(nums []int, target int) int {
	m := 0
	n := len(nums) - 1
	var doWork func(m, n int) int
	doWork = func(m, n int) int {

		mid := (m + n) / 2
		res := -1
		if target < nums[m] || target > nums[n] {
			return -1
		}
		if target == nums[m] {
			return m
		}
		if target == nums[n] {
			return n
		}
		if target == nums[mid] {
			return mid
		}
		if m+1 >= n {
			return -1
		}
		if target < nums[mid] {
			res = doWork(m, mid)
		} else {
			res = doWork(mid, n)
		}
		return res
	}
	return doWork(m, n)
}

//704. 二分查找
func search2(nums []int, target int) int {
	m, n := 0, len(nums)-1
	for m <= n {
		mid := (n-m)/2 + m
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			m = mid + 1
		} else {
			n = mid - 1
		}
	}
	return -1
}

/*
https://www.lintcode.com/problem/61/
61 · 搜索区间
给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。
如果目标值不在数组中，则返回[-1, -1]
in:	数组 = []
	target = 9
out: [-1,-1]

in:	数组 = [5, 7, 7, 8, 8, 10]
	target = 8
out: [3,4]
*/
func searchRange(A []int, target int) []int {
	notFind := []int{-1, -1}
	if len(A) == 0 {
		return notFind
	}

	left := 0
	right := len(A) - 1
	for left <= right {
		mid := left + (right-left)/2
		if A[mid] == target {
			// fmt.Println("==mid==", mid)
			// 分为两个
			if (mid == 0 && len(A) == 1) || (mid == len(A)-1) {
				return []int{mid, mid}
			}
			leftRes := searchRange(A[:mid+1], target)
			rightRes := searchRange(A[mid+1:], target)
			if equalSlice(leftRes, notFind) {
				return []int{rightRes[0] + mid + 1, rightRes[1] + mid + 1}
			}
			if equalSlice(rightRes, notFind) {
				return []int{leftRes[0], leftRes[1]}
			}

			return []int{leftRes[0], rightRes[1] + mid + 1}
		} else if target > A[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return notFind
}
func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if a == nil || b == nil {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func searchRange2(A []int, target int) []int {
	if len(A) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	start := 0
	end := len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] > target {
			end = mid
		} else if A[mid] < target {
			start = mid
		} else {
			// 如果相等，应该继续向左找，就能找到第一个目标值的位置
			end = mid
		}
	}
	// 搜索左边的索引
	if A[start] == target {
		result[0] = start
	} else if A[end] == target {
		result[0] = end
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	start = 0
	end = len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] > target {
			end = mid
		} else if A[mid] < target {
			start = mid
		} else {
			// 如果相等，应该继续向右找，就能找到最后一个目标值的位置
			start = mid
		}
	}
	// 搜索右边的索引
	if A[end] == target {
		result[1] = end
	} else if A[start] == target {
		result[1] = start
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	return result
}

//74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix) - 1
	// 0 1 2
	for left+1 < right {
		mid := left + (right-left)/2
		if matrix[mid][0] >= target {
			right = mid
		} else {
			left = mid
		}
	}

	var subLine []int
	if matrix[left][len(matrix[left])-1] >= target {
		subLine = matrix[left]
		leftSub := 0
		rightSub := len(subLine) - 1
		for leftSub <= rightSub {
			midSub := leftSub + (rightSub-leftSub)/2
			if subLine[midSub] == target {
				return true
			} else if subLine[midSub] > target {
				rightSub = midSub - 1
			} else {
				leftSub = midSub + 1
			}
		}
		return false
	} else {
		subLine = matrix[right]
		leftSub := 0
		rightSub := len(subLine) - 1
		for leftSub <= rightSub {
			midSub := leftSub + (rightSub-leftSub)/2
			if subLine[midSub] == target {
				return true
			} else if subLine[midSub] > target {
				rightSub = midSub - 1
			} else {
				leftSub = midSub + 1
			}
		}
		return false
	}
}

func isBadVersion(version int) bool {
	// do some thing
	return true
}

//278. 第一个错误的版本
func firstBadVersion(n int) int {

	left := 1
	right := n
	for left+1 < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if isBadVersion(left) {
		return left
	} else {
		return right
	}
}

//153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	// 6712345
	// 3456712
	// 1234567
	// 2345671
	// 21
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[left] {
			if nums[right] > nums[left] {
				return nums[left]
			}
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

//154. 寻找旋转排序数组中的最小值 II
func findMin2(nums []int) int {
	// 可重复
	// 22012
	// 20122
	// [3,3,1,3]
	left := 0
	right := len(nums) - 1
	for left+1 < right {
		mid := left + (right-left)/2

		fmt.Println("==left==", left, "==mid==", mid, " ==right==", right)
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left = left + 1
			right = right - 1
		} else if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] <= nums[right] {
		return nums[left]
	}
	return nums[right]
}

//33. 搜索旋转排序数组
func search3(nums []int, target int) int {
	/*
		给你 旋转后 的数组 nums 和一个整数 target，
		如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
	*/
	//输入：nums = [4,5,6,7,0,1,2], target = 0
	//输出：4
	// [1,3]
	// 思路：将数组在中间分成两部分，一定有一部分是有序数组

	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
	}

	res := -1
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	// 3 1
	if nums[left] < nums[mid] && nums[mid] < nums[right] {
		// 两部分都是有序的

		//1
		res = search2(nums[left:mid+1], target)
		if res != -1 {
			return res + left
		}
		//2
		res = search2(nums[mid+1:right+1], target)
		if res != -1 {
			return res + mid + 1
		}
	} else if nums[left] < nums[mid] {
		// 第一部分是有序
		res = search2(nums[left:mid+1], target)
		if res != -1 {
			return res + left
		}

		//第二部分
		res = search3(nums[mid+1:right+1], target)
		if res != -1 {
			return res + mid + 1
		}
	} else {
		//第二部分是有序
		res = search2(nums[mid:right+1], target)
		if res != -1 {
			return res + mid
		}
		// 第一部分
		res = search3(nums[left:mid], target)
		if res != -1 {
			return res + left
		}

	}

	return res
}

//81. 搜索旋转排序数组 II
func search4(nums []int, target int) bool {
	// nums可能重复，从小到大排序并旋转  存在为true、不存在为false
	//特点： 允许重复 需要考虑 left mid right 相等的情况，最坏的情况 全相等且不等于target O(N),平均是O(log(n))
	if nums == nil || len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}

	//开始二分
	left := 0
	right := len(nums) - 1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left += 1
			right -= 1
		} else if nums[mid] >= nums[left] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid
			}
		} else {
			if target <= nums[right] && target >= nums[mid] {
				left = mid
			} else {
				right = mid
			}
		}
	}
	if nums[left] == target || nums[right] == target {
		return true
	}
	return false
}

func main() {

	data := []int{5, 1, 3}

	// data := [][]int{{1, 3}}
	// fmt.Println("==1==", data[0:0], "==2==", data[0:1])

	res := search4(data, 3)
	fmt.Println("==res==", res)
}
