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
