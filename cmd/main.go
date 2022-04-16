package main

import (
	"fmt"
	"github.com/William-ZXS/algorithm"
)

func main() {
	data := []int{5, 1, 2, 4, 8, 3, 6, 7}
	//algorithm.QuickSort(data)
	data = algorithm.MergeSort(data)
	fmt.Println(data)
}
