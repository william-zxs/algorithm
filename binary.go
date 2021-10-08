package main

import "fmt"

//136. 只出现一次的数字
func singleNumber1(nums []int) int {
	// 使用hash表的方式  时间复杂度和空间复杂度都是O(n)
	res := 0
	data := make(map[int]int, 0)

	for _, num := range nums {
		if _, ok := data[num]; ok {
			delete(data, num)
		} else {
			data[num] += 1
		}

	}
	for k := range data {
		res = k
	}
	return res
}

//136. 只出现一次的数字
func singleNumber(nums []int) int {
	// 异或
	// 时间复杂度 O(n) 空间复杂度O(1)
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}

func main() {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	fmt.Println("==res==", res)

	data := make(map[int]int, 0)
	data[1] = 1
	data[2] = 2
	delete(data, 1)
	fmt.Println("==data==", data)
}
