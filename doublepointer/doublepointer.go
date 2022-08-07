package doublepointer

import (
	"github.com/William-ZXS/algorithm/util"
	"math"
)

//双指针 和 滑动窗口 类型题目

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

//76. 最小覆盖子串
func minWindow(s string, t string) string {
	sMap := make(map[byte]int, 0)
	tMap := make(map[byte]int, 0)

	for i := 0; i < len(t); i++ {
		tMap[t[i]] += 1
	}

	minStr := ""

	j := 0
	i := 0
	for i <= len(s) && j <= i {
		if !checkRes(sMap, tMap) {
			if i >= len(s) {
				break
			}

			//右指针移动
			if _, ok := tMap[s[i]]; ok {
				sMap[s[i]] += 1
			}
			i++

			continue
		}
		minStr = getMinStr(minStr, s[j:i])
		//左指针移动
		if _, ok := tMap[s[j]]; ok {
			sMap[s[j]] -= 1
		}

		j++

	}
	return minStr

}

func getMinStr(minStr, curStr string) string {
	if minStr == "" {
		return curStr
	}

	if len(minStr) > len(curStr) {
		return curStr
	}
	return minStr
}

func checkRes(sMap, tMap map[byte]int) bool {
	for k, v := range tMap {
		if sMap[k] < v {
			return false
		}
	}
	return true
}

//567. 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var array1 [26]int
	var array2 [26]int
	for i, v := range s1 {
		array1[v-'a'] += 1
		array2[s2[i]-'a'] += 1
	}
	if array1 == array2 {
		return true
	}

	for i, j := 0, len(s1); j < len(s2); i, j = i+1, j+1 {
		//左指针移动
		array2[s2[i]-'a'] -= 1
		//右指针移动
		array2[s2[j]-'a'] += 1
		if array1 == array2 {
			return true
		}

	}
	return false

}
