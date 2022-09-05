package main

import "fmt"

//fmt.Println 函数的参数是 interface。
//对于内置类型，函数内部会用穷举法，得出它的真实类型，然后转换为字符串打印。
//而对于自定义类型，首先确定该类型是否实现了 String() 方法，如果实现了，则直接打印输出 String() 方法的结果；
//否则，会通过反射来遍历对象的成员进行打印。

//类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖。
type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

func main() {
	var s = Student{
		Name: "qcrao",
		Age:  18,
	}

	fmt.Println(s)
}
