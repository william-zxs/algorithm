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

func main() {
	res := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println("==res==:", res)
}
