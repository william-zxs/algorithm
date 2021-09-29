# algorithm
algorithm practice



## 算法
* DFS

（Deep First Search）深度优先搜索
* BFS

（Breath First Search）广度优先搜索


* 中序遍历
中序遍历（LDR）是二叉树遍历的一种，也叫做中根遍历、中序周游。在二叉树中，中序遍历首先遍历左子树，然后访问根结点，最后遍历右子树。



## 计算时间复杂度


## 链表
当链表头要变动的时候，需要使用一个dummyNode(哑巴节点)


## golang
slice作为参数传递给函数


* slice的头部插入、指定位置的插入
```
func insert(nums []list,index int,num int){
    nums = append(nums,0)
    copy(nums[index+1:],nums[index:])
    nums[index] = num
}
```

* 反转slice
```
func reverse(nums [][]int) {
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
```

* 切片


```
//切片解包
data1 := []int{1,2,3}
data2 := []int{1,2,3}
data1 = append(data1,data2...)
```
```
//这样操作是可以的，结果是空的切片
data := []int{1}
data = data[1:]
```

* 字符串
```
str := "Hello,世界"
//utf-8遍历
for i := 0; i < len(str); i++ {
    ch := str[i]
    fmt.Println(ch)
}
fmt.Println("=============>Unicode遍历")
//Unicode遍历
for _, ch1 := range str {
    fmt.Println(ch1)
}
```
```
1、双引号里的字符串可以转义，不能换行

2、反引号里面的内容不能转义，可以换行，一般用于SQL语句，html等大段内容，以及正则表达式的使用

3、单引号，一般只能用来包裹一个字节的ASCII码字符，例如：
var asc byte = ‘a’
fmt.Println(asc) //输出a的ASCII码值 97
```

### TODO

golang切片的指针