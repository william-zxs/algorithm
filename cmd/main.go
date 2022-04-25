package main

import (
	"fmt"
	"github.com/William-ZXS/algorithm"
)

type user struct {
	Name string
	Age  int
}

func (u user) String() string {
	return "aaaaaaaa"
}

func main() {
	data := []int{5, 1, 2, 4, 8, 3, 6, 7}
	//algorithm.QuickSort(data)
	//data = algorithm.MergeSort(data)
	//data = algorithm.HeapSort(data)
	data = algorithm.BubblingSort(data)
	fmt.Println(data)

}
