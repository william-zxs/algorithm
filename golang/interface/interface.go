package main

import (
	"fmt"
)

//类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖。
//所以， Student 结构体定义了接受者类型是值类型的 String() 方法时，通过
//
//fmt.Println(s)
//fmt.Println(&s)
//均可以按照自定义的格式来打印。
//
//如果 Student 结构体定义了接受者类型是指针类型的 String() 方法时，只有通过
//
//fmt.Println(&s)
//才能按照自定义的格式打印。

//使用过如下的方式检查自己的定义的类型是否实现了某个接口
//var _ io.Writer = (*myWriter)(nil)

type Coder interface {
	code()
}

type Gopher struct {
	name string
}

func (r Gopher) code() {
	fmt.Println(r.name)
}

func testInterface(c Coder) {
	fmt.Printf("%T %v\n", c, c)
	c.code()

	//类型断言
	switch a := c.(type) {
	case Gopher:
		fmt.Println(a.name)
	}
}

type Person interface {
	growUp()
}

type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}

func main() {
	//var qcrao = Person(Student{age: 18})
	//
	//fmt.Println(qcrao)

	a := Gopher{
		name: "i am a gopher",
	}
	testInterface(a)
}
