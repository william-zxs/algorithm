package main

import (
	"fmt"
	"sync"
)

func ChanSlice() {
	var ch chan int
	ch = make(chan int, 1)

	s := make([]int, 0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			//for v := range ch {
			//}
			ch <- num
			s = append(s, num)
			v, ok := <-ch
			fmt.Println("==receive==", v, " ==ok==", ok)

			wg.Done()
		}(i)

	}
	wg.Wait()
	close(ch)
	fmt.Println("==close==")
	fmt.Println("==len==", len(s))
}

func PlainSlice() {
	s := make([]int, 0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			s = append(s, num)
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println("==len==", len(s))
}
