package dfs

//已知一个无序数组 array，元素均为正整数。
//给定一个目标值 target，输出数组中是否存在若干元素的组合，相加为目标值。
//输入：array=[2，1，2，5，10，10，20，50], target=71
//输出：true
//输入：array=[2，1，2，5，10，10，20，50], target=101
//输出：false
func TargetSum(array []int, target int) bool {
	for i := 0; i < len(array); i++ {
		v := target
		if array[i] > v {
			continue
		}
		if array[i] == v {
			return true
		} else if array[i] < v {
			v = v - array[i]
			if TargetSum(array[i+1:], v) {
				return true
			}
		}
	}
	return false
}
