package main

import (
	"context"
	"fmt"
	"time"
)

func Vctx(ctx context.Context) {
	v, ok := ctx.Value("name").(string)
	fmt.Println(v, ok)
	ctx = context.WithValue(ctx, "age", 20)
	Cctx(ctx)

}

func Cctx(ctx context.Context) {
	v, ok := ctx.Value("name").(string)
	v2, ok2 := ctx.Value("age").(int)
	fmt.Println("==1==", v, ok)
	fmt.Println("==2==", v2, ok2)
	fmt.Println("==ctx==", ctx)

}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "william")
	//ctx, cancel := context.WithCancel(ctx)
	go Vctx(ctx)
	//cancel()
	fmt.Println("====")
	fmt.Println(<-ctx.Done())
	time.Sleep(time.Second)
	//timer := time.NewTimer(10)

}
