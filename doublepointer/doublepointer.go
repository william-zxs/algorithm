package doublepointer

import (
	"github.com/William-ZXS/algorithm/util"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

//438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var pArray [26]int
	var sArray [26]int
	for i, v := range p {
		pArray[v-'a'] += 1
		sArray[s[i]-'a'] += 1
	}
	res := make([]int, 0)
	if pArray == sArray {
		res = append(res, 0)
	}
	for i, j := 0, len(p); j < len(s); i, j = i+1, j+1 {
		sArray[s[i]-'a'] -= 1
		sArray[s[j]-'a'] += 1
		if pArray == sArray {
			res = append(res, i+1)
		}
	}
	return res

}

//3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxL := 0
	state := make(map[byte]int, 0)

	for r, l := 0, 0; r < len(s); r++ {
		state[s[r]]++
		if state[s[r]] == 1 {
			maxL = max(maxL, r-l+1)
		} else {
			for l < r {
				state[s[l]]--
				if s[l] == s[r] {
					l++
					break
				}
				l++
			}
		}
	}
	return maxL
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

//19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p1 := head
	i := 1
	for ; i < n; i++ {
		p1 = p1.Next
	}
	p2 := head
	pre := dummy
	for p1.Next != nil {
		p1 = p1.Next
		pre = p2
		p2 = p2.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
