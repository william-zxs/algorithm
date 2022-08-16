package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func mutexAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a += 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use mutex a is %d, spend time: %v\n", a, timeSpends)
}

func AtomicAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use atomic a is %d, spend time: %v\n", atomic.LoadInt32(&a), timeSpends)
}

type Rectangle struct {
	length int
	width  int
}

var rect atomic.Value

func update(width, length int) {
	rectLocal := new(Rectangle)
	rectLocal.width = width
	rectLocal.length = length
	rect.Store(rectLocal)
}

func testvalue() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	// 10 个协程并发更新
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			update(i, i+5)
		}()
	}
	wg.Wait()
	_r := rect.Load().(*Rectangle)
	fmt.Printf("rect.width=%d\nrect.length=%d\n", _r.width, _r.length)
}

func main() {

}
