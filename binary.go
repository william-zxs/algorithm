package main

import "fmt"

/*
	常见的位逻辑运算符:
　　　　&:
　　　　　　位与运算符,表示AND(表示所有条件都得匹配),运算规则为相同位都是1时结果才为1，不同则为0。举个例子:如"5 & 7",结果为5。
　　　　　　计算过程为:我们用一个字节来表示一个数字，5转换成二进制为0000 0101,7转换成二进制为0000 0111,此时做位与运算，相同位都是1时结果才为1，最终得到结果二进制结果0000 0101，使用十进制表示为5。

　　　　|:
　　　　　　位或运算符,表示OR(表示有一个条件匹配即可),运算规则为相同位只要一个为1则为1。举个例子:如"5 | 7",结果为7。
　　　　　　计算过程为:我们用一个字节来表示一个数字，5转换成二进制为0000 0101,7转换成二进制为0000 0111,此时做位或运算，相同位只要一个为1时结果才为1，最终得到结果0000 0111。使用十进制表示为7。

　　　　^:
　　　　　　位异或运算符,表示XOR,运算规则为相同位不同则为1，相同则为0。举个例子:如"5 ^ 7",结果为2。
　　　　　　计算过程为:我们用一个字节来表示一个数字，5转换成二进制为0000 0101,7转换成二进制为0000 0111,此时做位异或运算，相同位不同则为1，相同则为0，最终得到结果0000 0010。使用十进制表示为2。
　　　　&^:
　　　　　　位清空运算符，表示AND NOT，运算规则为后数为0，则用前数对应位代替，后数为1则取0。举个例子:如"5 ^ 7",结果为0。
　　　　　　计算过程为:我们用一个字节来表示一个数字，5转换成二进制为0000 0101,7转换成二进制为0000 0111,此时做位清空运算符，为后数为0，则用前数对应位代替，后数为1则取0，最终得到结果0000 0000。使用十进制表示为0。
	常见的位移运算符:
　　　　<<:
　　　　　　左移，表示将对应的二进制数字向左移动相应的位数，比如 5 << 3,结果为40。
　　　　　　计算过程为:我们用一个字节来表示一个数字，5转换成二进制为0000 0101,将二进制数字向左位移3位得到0010 1000,使用十进制表示为"40"。
　　　　>>:
　　　　　　右移，表示将对应的二进制数字向右移动相应的位数，比如 5 >> 3,结果为1。
　　　　　　计算过程为:我们用一个字节来表示一个数字，5转换成二进制为0000 0101,将二进制数字向左位移3位得到0000 0001,使用十进制表示为"1"。
*/

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

//137. 只出现一次的数字 II
func singleNumber22(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			sum += num >> i & 1
		}
		if sum%3 > 0 {
			ans = ans | (1 << i)
		}
	}
	return int(ans)
}

//260. 只出现一次的数字 III
func singleNumber3(nums []int) []int {
	// 解题思路
	// 通过异或的结果 找到一个1的位，说明两个出现的一次的数这位是不同的，所以用来分成两组，再异或求值就可以了
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}

	index := 0
	for i := 0; i < 32; i++ {
		diff := res >> i & 1
		if diff == 1 {
			index = i
			break
		}
	}
	numOne := 0
	numTwo := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]>>index&1 == 0 {
			numOne = numOne ^ nums[i]
		} else {
			numTwo = numTwo ^ nums[i]
		}

	}
	return []int{numOne, numTwo}
}

//191. 位1的个数
func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if num>>i&1 == 1 {
			count += 1
		}
	}
	return count
}

//191. 位1的个数
func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// 338. 比特位计数
func countBits(n int) []int {
	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		num := i
		count := 0
		for num != 0 {
			num = num & (num - 1)
			count++
		}
		res = append(res, count)
	}
	return res
}

// 338. 比特位计数
func countBits2(n int) []int {
	//动态规划的方式
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

//190. 颠倒二进制位
func reverseBits(num uint32) uint32 {
	//思路
	//通过位移运算符计算每一位，然后位移到正确位置，相加得到最终结果
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res += (num >> i & 1) << (31 - i)
	}
	return uint32(res)
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	res := singleNumber3(nums)
	fmt.Println("==res==", res)
}
