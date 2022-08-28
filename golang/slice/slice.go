package main

import (
	"fmt"
	"unsafe"
)

//slice 结构
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func grow() {
	//	扩容和底层数组特性
	s := []int{5}
	s = append(s, 6)
	s = append(s, 7)

	x := append(s, 8)
	y := append(s, 9)
	fmt.Println(s, x, y)
}

func TestA() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 2,3,4 len
	// {2, 3, 4, 5, 6, 7, 8, 9} cap
	s1 := s[2:5]
	fmt.Println("==len==", len(s1))
	fmt.Println("==cap==", cap(s1))

	// 4,5,6,7 len
	// 4,5,6,7,8 cap
	s2 := s1[2:6:7]
	//
	fmt.Println("==s2==", len(s2))
	s2 = append(s2, 100)
	fmt.Println("==slice==", s)
	s2 = append(s2, 200)
	fmt.Println("==s2==", s2)
	//
	//s1[2] = 20

	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(slice)
}

func main() {

}
