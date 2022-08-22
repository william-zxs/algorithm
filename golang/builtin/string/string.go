package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//了解了 string 类型的实现原理后，我们还可以得到这样一个结论，
//那就是我们直接将 string 类型通过函数 / 方法参数传入也不会带来太多的开销。
//因为传入的仅仅是一个“描述符”，而不是真正的字符串数据。

//不可改变： “string类型虽然是不能更改的，但是可以被替换，
//因为stringStruct中的str指针是可以改变的，
//只是指针指向的内容是不可以改变的，也就说每一次更改字符串，
//就需要重新分配一次内存，之前分配的空间会被gc回收。”

//引号：单引号、双引号、反引号
//单引号：只能表示一个rune字节（一个Unicode),不能表示多个字符,Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
//双引号：双引号用来创建可解析的字符串，支持转义
//反引号：反引号用来创建原生的字符串字面量，可能由多行组成，但不支持转义，并且可以包含除了反引号外其他所有字符
//双引号创建可解析的字符串应用最广泛，反引号用来创建原生的字符串则多用于书写多行消息，HTML以及正则表达式。

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func changeStrDataByUnsafe() {
	s1 := "aaa"
	s2 := "ccc"
	hdr1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	hdr2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("==s1== %v %p %v\n", s1, &s1, hdr1.Data)
	fmt.Printf("==s2== %v %p %v\n", s2, &s2, hdr2.Data)
	hdr1.Data = hdr2.Data
	hdr1.Len = hdr2.Len
	fmt.Printf("==s1== %v %p %v\n", s1, &s1, hdr1.Data)

}

//字符串好似不可变的
func unchangeStr() {
	//字符串不可变
	var s string = "hello"
	//s[0] = 'k'   // 错误：字符串的内容是不可改变的
	s = "gopher" // ok
	fmt.Println(s)
}

func main() {
	changeStrDataByUnsafe()
}
