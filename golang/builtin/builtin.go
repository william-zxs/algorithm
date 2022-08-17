package main

import (
	"errors"
	"fmt"
)

//查看汇编
//go tool compile -N -l -S builtin.go

const (
	one = iota
	two
	three = 4
	four  = iota
	five
)

func testIota() {
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
	fmt.Println(five)
}

//单引号在go语言中表示golang中的rune(int32)类型，单引号里面是单个字符

//在go语言中双引号里面可以是单个字符也可以是字符串，双引号里面可以有转义字符，如\n、\r等，对应go语言中的string类型。

//反引号
//反引号中的字符表示其原生的意思，在单引号中的内容可以是多行内容，不支持转义

//defer
//defer 与 return 的返回时机
//这里我先说结论，总结一下就是，函数的整个返回过程应该是：
//1. return 对返回变量赋值，如果是匿名返回值就先声明再赋值；
//2. 执行 defer 函数；
//3. return 携带返回值返回。

//1.A deferred function’s arguments are evaluated when the defer statement is evaluated.
//
//2.Deferred function calls are executed in Last In First Out order after the surrounding function returns.
//
//3.Deferred functions may read and assign to the returning function’s named return values.

//函数return步骤
//函数在返回时，首先函数返回时会自动创建一个返回变量假设为ret(如果是命名返回值的函数则不会创建)，函数返回时要将变量i赋值给ret，即有ret = i。
//然后检查函数中是否有defer存在，若有则执行defer中部分。
//最后返回ret
func Anonymous() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2 value is ", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1 in value is ", i)
	}()
	i++
	return i
}

func HasName() (j int) {
	defer func() {
		j++
		fmt.Println("defer2 in value", j)
	}()

	defer func() {
		j++
		fmt.Println("defer1 in value", j)
	}()

	return j
}

// 测试1
func Test1() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}

func Test2() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)
	return 2
}

func Test3() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
}

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
}

func TestDefer() {
	fmt.Println(Test1())
	fmt.Println(Test2())
	fmt.Println(Test3())

	e1()
	e2()
	e3()
}
func main() {

	defer func() {
		defer func() {
			inerr := recover()
			fmt.Println("inerr", inerr)
		}()
		fmt.Println("00000")
		d := recover()
		fmt.Println(d)
		fmt.Println("11111")
		panic("bbbb")
		fmt.Println("22222")

		fmt.Println("333333")
	}()
	a := 10
	if a > 5 {
		panic("aaa")
	}
}
