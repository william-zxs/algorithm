package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 2,3,4 len
	// {2, 3, 4, 5, 6, 7, 8, 9} cap
	s1 := slice[2:5]
	fmt.Println("==len==", len(s1))
	fmt.Println("==cap==", cap(s1))

	// 4,5,6,7 len
	// 4,5,6,7,8 cap
	s2 := s1[2:6:7]
	//
	fmt.Println("==s2==", len(s2))
	s2 = append(s2, 100)
	fmt.Println("==slice==", slice)
	s2 = append(s2, 200)
	fmt.Println("==s2==", s2)
	//
	//s1[2] = 20

	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(slice)
}
