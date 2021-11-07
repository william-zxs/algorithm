package main

import (
	"fmt"
	"math"
)

/*
滑动窗口思想
*/

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	//两个 map 一个记录目标  一个记录当前状态
	targetMap := make(map[byte]int, 0)
	currentMap := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}
	mactch := 0
	sLen := math.MaxInt32
	sL := -1
	sR := -1

	for l, r := 0, 0; r < len(s); r++ {
		//  currentMap[s[r]]
		if count, ok := targetMap[s[r]]; ok && count > 0 {
			currentMap[s[r]]++
			if currentMap[s[r]] == targetMap[s[r]] {
				mactch++
			}
		}
		fmt.Println("==l==:", l, "==r==:", r, "==mactch==:", mactch)
		for mactch == len(targetMap) && l <= r {

			// 计算长度，保存位置
			if r-l+1 < sLen {

				sLen = r - l + 1
				sL = l
				sR = r + 1
				fmt.Println("==sLen==", sLen)
			}

			if currentMap[s[l]] == targetMap[s[l]] && targetMap[s[l]] != 0 {
				mactch--
			}
			currentMap[s[l]]--
			l++
			fmt.Println("==mactch==", mactch)
		}
	}
	if sL == -1 {
		return ""
	}

	return s[sL:sR]
}

// 567. 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	targetMap := make(map[byte]int, 0)
	currentMap := make(map[byte]int, 0)
	matchCount := 0

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		targetMap[s1[i]]++
		currentMap[s2[i]]++
	}
	for k, v := range targetMap {
		if currentMap[k] == v {
			matchCount++
		}
	}
	if len(targetMap) == matchCount {
		return true
	}

	//滑动窗口
	for r := len(s1); r < len(s2); r++ {

		//左边
		if targetMap[s2[r-len(s1)]] > 0 {
			currentMap[s2[r-len(s1)]]--
			if currentMap[s2[r-len(s1)]] == targetMap[s2[r-len(s1)]] {
				matchCount++
			} else if currentMap[s2[r-len(s1)]] == targetMap[s2[r-len(s1)]]-1 {
				matchCount--
			}
		}

		//右边
		if targetMap[s2[r]] > 0 {
			currentMap[s2[r]]++
			if currentMap[s2[r]] == targetMap[s2[r]] {
				matchCount++
			} else if currentMap[s2[r]] == targetMap[s2[r]]+1 {
				matchCount--
			}
		}
		if len(targetMap) == matchCount {
			return true
		}
	}

	return false
}

func main() {
	// res := minWindow("ADOBECODEBANC", "ABC")
	res := checkInclusion("adc", "dcda")
	fmt.Println("==res==:", res)
}
