# algorithm
algorithm practice

## 数据结构



## 树
**种类**

红黑树、区间树、B树

* 二叉搜索树
二叉搜索树关键的性质是根节点的值大于左子树所有节点的值，小于右子树所有节点的值，且左子树和右子树也同样为二叉搜索树
* AVL树
  AVL树是具有以下性质的二叉搜索树：
  它的左右子树都是AVL树
  左右子树高度之差(简称平衡因子)的绝对值不超过1(-1/0/1)
  如果一棵二叉搜索树是高度平衡的，它就是AVL树。
* 平衡二叉树
* 完全二叉树
是一种特殊的二叉树，满足以下要求：
所有叶子节点都出现在 k 或者 k-1 层，而且从 1 到 k-1 层必须达到最大节点数；
第 k 层可以不是满的，但是第 k 层的所有节点必须集中在最左边。
需要注意的是不要把完全二叉树和“满二叉树”搞混了，完全二叉树不要求所有树都有左右子树，但它要求：
任何一个节点不能只有右子树没有左子树
叶子节点出现在最后一层或者倒数第二层，不能再往上
* 满二叉树

* 堆
「堆」是一个完全二叉树

## 算法
* DFS

（Deep First Search）深度优先搜索
* BFS

（Breath First Search）广度优先搜索


* 中序遍历
中序遍历（LDR）是二叉树遍历的一种，也叫做中根遍历、中序周游。在二叉树中，中序遍历首先遍历左子树，然后访问根结点，最后遍历右子树。二叉搜索树中序遍历的结果是升序的。

* 动态规划

* 贪心算法

* 回溯

## 排序
* 快排

## 计算时间复杂度


## 链表
当链表头要变动的时候，需要使用一个dummyNode(哑巴节点)


## golang
slice作为参数传递给函数，是值传递，里面的是值的地址拷贝，修改slice中的值是可以生效的，如果添加值，则不会影响到外部slice


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
var asc byte = 'a'
fmt.Println(asc) //输出a的ASCII码值 97
```

## 二进制
n & (n−1)，其运算结果恰为把 nn 的二进制位中的最低位的 11 变为 00 之后的结果

汉明重量、汉明距离

## 位运算

### TODO

golang切片的指针