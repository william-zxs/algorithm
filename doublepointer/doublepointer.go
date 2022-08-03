package doublepointer

import (
	"github.com/William-ZXS/algorithm/util"
	"math"
)

//双指针
//当两个指针同速同方向移动，也就是滑动窗口
func MaxSum(cycle []int, N int) int {
	sum := math.MinInt
	for i, j := 0, N-1; j < len(cycle)*2-1; i, j = i+1, j+1 {
		v := 0
		for m := i; m <= j; m++ {
			v += cycle[m%len(cycle)]
		}
		sum = util.Max(sum, v)
	}
	return sum
}

//344. 反转字符串
func reverseString(s []byte) []byte {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}
