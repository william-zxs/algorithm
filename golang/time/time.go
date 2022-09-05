package main

//ticker定时器表示每隔一段时间就执行一次，一般可执行多次。
//timer定时器表示在一段时间后执行，默认情况下只执行一次，如果想再次执行的话，每次都需要调用 time.Reset() 方法，此时效果类似ticker定时器。同时也可以调用 Stop() 方法取消定时器
//timer定时器比ticker定时器多一个 Reset() 方法，两者都有 Stop() 方法，表示停止定时器,底层都调用了stopTimer() 函数。
//除了上面的定时器外，Go 里的 time.Sleep 也起到了类似一次性使用的定时功能。只不过 time.Sleep 使用了系统调用。而像上面的定时器更多的是靠 Go 的调度行为来实现。
//无论哪种计时器，.C 都是一个 chan Time 类型且容量为 1 的单向 Channel，当有超过 1 个数据的时候便会被阻塞，以此保证不会被触发多次。

//Timer 允许再次被启用，而 time.After 返回的是一个 channel，将不可复用。
//而且需要注意的是 time.After 本质上是创建了一个新的 Timer 结构体，只不过暴露出去的是结构体里的 channel 字段而已。
//因此如果在 for{...}里循环使用了 time.After，将会不断的创建 Timer。如下的使用方法就会带来性能问题
func main() {

}
