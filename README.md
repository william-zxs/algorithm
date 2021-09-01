# algorithm
algorithm learn


## 计算时间复杂度




## golang
slice作为参数传递给函数


* slice的头部插入、指定位置的插入
func insert(nums []list,index int,num int){
    nums = append(nums,0)
    copy(nums[index+1:],nums[index:])
    nums[index] = num
}


* 反转slice
func reverse(nums [][]int) {
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
