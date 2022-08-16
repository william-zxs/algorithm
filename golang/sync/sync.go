package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var V string

func getV() string {
	once.Do(
		func() {
			V = "william"
		})
	return V
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(getV())
	}
}
