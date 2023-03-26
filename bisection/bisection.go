package bisection

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 中等
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	res[0] = findLeft(nums, target)
	if res[0] == -1 {
		return res
	}
	res[1] = findRight(nums, target)
	if res[1] == -1 {
		res[1] = res[0]
	}
	return res
}
func findLeft(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}
func findRight(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		}
	}
	if nums[l-1] == target {
		return l - 1
	}
	return -1
}
