package main

import (
	"fmt"
	"time"
)

//Channel 可能会引发 goroutine 泄漏。
//泄漏的原因是 goroutine 操作 channel 后，处于发送或接收阻塞状态，而 channel 处于满或空的状态，一直得不到改变。
//同时，垃圾回收器也不会回收此类资源，进而导致 gouroutine 会一直处于等待队列中，不见天日。
//另外，程序运行过程中，对于一个 channel，如果没有任何 goroutine 引用了，gc 会对其进行回收操作，不会引起内存泄漏。

//在 Go 语言中，对于一个 channel，如果最终没有任何 goroutine 引用它，不管 channel 有没有被关闭，
//最终都会被 gc 回收。所以，在这种情形下，所谓的优雅地关闭 channel 就是不关闭 channel，让 gc 代劳。

//range 会阻塞的形式接受数据，直到channel close 会往下执行
//for data := range ch {
//
//}

//默认在select中break是只跳脱了select体，而不是结束for循环
func Break(ch <-chan int) {

LOOP:
	for {
		select {
		case a, ok := <-ch:
			if !ok {
				fmt.Println("==break==")
				//这里break只是会跳出select，如果需要跳出for循环，
				//需要 break LOOP，或者直接return
				break LOOP
			}
			fmt.Println(a, ok)
		default:
			fmt.Println("default")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

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

func deadlock() {
	//这种情况会死锁
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	for data := range ch {
		fmt.Println(data)
	}
}
func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch <- 10
	}()

	select {
	case a := <-ch:
		print(a)
	default:
		print("default")
	}
	
}
