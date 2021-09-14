# algorithm
algorithm practice



## 算法
* DFS

（Deep First Search）深度优先搜索
* BFS

（Breath First Search）广度优先搜索

## 计算时间复杂度


## 链表
当链表头要变动的时候，需要使用一个dummyNode(哑巴节点)


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
