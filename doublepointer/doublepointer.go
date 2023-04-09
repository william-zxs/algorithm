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

// 344. 反转字符串
func reverseString(s []byte) []byte {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}

// 76. 最小覆盖子串
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

// 567. 字符串的排列
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

// 438. 找到字符串中所有字母异位词
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

// 3. 无重复字符的最长子串
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

// 19. 删除链表的倒数第 N 个结点
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

// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {

	//快慢双指针
	i := head
	j := head
	for i != nil && i.Next != nil {
		i = i.Next.Next
		j = j.Next
	}
	return j
}

// 判断链表是否包含环
func HasRing(head *ListNode) bool {
	//快慢双指针，相遇则有环，如果快指针遇到nil，则无环
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if f == s {
			return true
		}
	}
	return false
}

// 如果链表中含有环，如何计算这个环的起点？
func DetectCycle(head *ListNode) *ListNode {
	//快慢指针，如果快指针遇到nil，则没有环；如果快慢指针相遇，则有环;
	//此时慢指针回到head，快指针和慢指针同速度，相遇点即为环起点
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			//说明有环
			break
		}
	}
	if f == nil || f.Next == nil {
		//说明无环
		return nil
	}

	s = head
	for s != f {
		s = s.Next
		f = f.Next
	}
	return s
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

// 167. 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l = l + 1
		} else {
			r = r - 1
		}
	}
	return []int{-1, -1}
}

// 344. 反转字符串
func reverseString2(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 5. 最长回文子串
func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var res string
	for i := 0; i < len(s)-1; i++ {
		res1 := helper(s, i, i)
		res2 := helper(s, i, i+1)
		if len(res) < len(res1) {
			res = res1
		}
		if len(res) < len(res2) {
			res = res2
		}

	}
	return res
}

func helper(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	l = l + 1
	r = r
	if l > r {
		return ""
	}
	return s[l:r]
}

// 11. 盛最多水的容器
// 中等
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	var max int
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}

// 264. 丑数 II
func nthUglyNumber(n int) int {
	//合并三个slice
	n1, n2, n3 := 2, 3, 5
	p1, p2, p3 := 0, 0, 0
	res := []int{1}
	for len(res) < n {
		v1 := res[p1] * n1
		v2 := res[p2] * n2
		v3 := res[p3] * n3
		v := min(min(v1, v2), v3)
		if v > res[len(res)-1] {
			res = append(res, v)
		}
		switch v {
		case v1:
			p1 += 1

		case v2:
			p2 += 1

		case v3:
			p3 += 1

		}

	}
	return res[len(res)-1]
}
