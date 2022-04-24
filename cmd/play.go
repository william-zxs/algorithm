package main

import (
	"fmt"
	"unsafe"
)

func Play() {
	a := "abcd"
	fmt.Printf("==%v==%+v \n", a, a)
	u := user{}
	fmt.Println(unsafe.Sizeof(u))
}
