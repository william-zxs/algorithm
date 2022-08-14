package backtrack

//回溯法

//46. 全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(nums []int, track []int, used []bool)
	backtrack = func(nums []int, track []int, used []bool) {
		if len(track) == len(nums) {
			result = append(result, append([]int{}, track...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtrack(nums, track, used)
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backtrack(nums, track, used)

	return result
}
