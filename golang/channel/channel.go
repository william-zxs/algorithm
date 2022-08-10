package main

import (
	"fmt"
	"time"
)

// timer ticker

func testclose() {
	ch := make(chan int)
	go func() {
		for {
			v, ok := <-ch
			if ok {
				fmt.Println("==ok==", v, time.Now())
			} else {
				fmt.Println("==not ok==", v, time.Now())
			}
			time.Sleep(time.Millisecond * 500)
		}

	}()
	time.Sleep(time.Second)
	ch <- 10
	time.Sleep(time.Second * 2)
	close(ch)
	//ch <- 10
	fmt.Println()
	time.Sleep(time.Second * 2)
}

func testTimer() {
	t := time.NewTimer(time.Second * 2)
	ch := make(chan bool)
	go func() {
		for {
			select {
			case d := <-t.C:
				fmt.Println("==t.C value==", d, time.Now())
				t.Reset(time.Second)
			case s := <-ch:
				fmt.Println("==s==", s)
				if s {
					fmt.Println("==return==")
					return
				}
			}
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("==stop==")
	t.Stop()
	//ch <- false
	time.Sleep(time.Second * 5)
	ch <- true
	time.Sleep(time.Second)
}

func main() {
	testTimer()
}
