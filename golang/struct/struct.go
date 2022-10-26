package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	F1 int64
	F2 int64
	F3 int8
	F4 int16
}

//空 struct的 坑点
func compareZeroStruct() {
	//https://mp.weixin.qq.com/s/K5B2ItkzOb4eCFLxZI5Wvw
	a := new(struct{})
	b := new(struct{})
	//打开print 则下面的比较相等，因为print会使其逃逸，地址都是zerobase，go run -gcflags="-m -l" playground.go
	//注释掉就不相等了，变量没有逃逸，编译器会直接返回false
	//fmt.Println(a, b)
	fmt.Println("a == b:", a == b)
}

func main() {
	s := S{
		F1: 1,
		F2: 2,
		F3: 3,
		F4: 4,
	}
	//fmt.Println(unsafe.Sizeof(s))
	//fmt.Println(unsafe.Offsetof(s.F1))
	//fmt.Println(unsafe.Offsetof(s.F2))
	//fmt.Println(unsafe.Offsetof(s.F3))
	//fmt.Println(unsafe.Offsetof(s.F4))
	fmt.Println(unsafe.Alignof(s.F4))
	a := []int{}
	fmt.Println(unsafe.Sizeof(a))

}
