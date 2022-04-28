package main

import (
	"fmt"
	"unsafe"
)

type AA struct{}

type BB struct {
	AA
}

func work(data AA) {
	fmt.Println("====")
}

func Play() {
	a := "abcd"
	fmt.Printf("==%v==%+v \n", a, a)
	u := user{}
	fmt.Println(unsafe.Sizeof(u))
}
