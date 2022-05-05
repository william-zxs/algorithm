package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type user struct {
	Name string
	Age  int
}

func (u user) String() string {
	return "aaaaaaaa"
}

type Goose struct {
	age  int
	name string
}

type Bird interface {
	fly()
	sing()
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
	fmt.Println(x)
}

func testFoo() {
	var x *int = nil
	Foo(x)
}
func testRune() {
	a := '你'
	fmt.Println(a)

}

func testEmptySlice() {
	var s1 []int         // nil
	s2 := make([]int, 0) //empty
	s3 := make([]int, 0)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s3)
}

func change(data *[5]int) {
	fmt.Println("==*arr1==", unsafe.Sizeof(data))

	data[0] = 10
}

type J struct {
	a string //小写无tag
	b string `json:"B"` //小写+tag
	C string //大写无tag
	D string `json:"DD"` //大写+tag
}

func jsonTest() {
	j := J{
		a: "1",
		b: "2",
		C: "3",
		D: "4",
	}
	fmt.Printf("转为json前j结构体的内容 = %+v\n", j)
	jsonInfo, _ := json.Marshal(j)
	fmt.Printf("转为json后的内容 = %+v\n", string(jsonInfo))

	jsonStr := []byte(`{"B":"5","C":"10","DD":"20"}`)
	json.Unmarshal(jsonStr, &j)
	fmt.Println("==j==", j)

}

type Cat struct {
	name string
	age  int
}

func compare() {
	cat1 := Cat{"a", 10}
	cat2 := Cat{"a", 10}
	fmt.Println(cat1 == cat2)
	fmt.Printf("%p  %p\n", &cat1, &cat2)
}

func testSync() {
	var sum uint64 = 0
	var sum2 uint64 = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum++
		}()
	}
	wg.Wait()
	fmt.Println("==sum==", sum)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&sum2, 1)
		}()
	}
	wg.Wait()
	fmt.Println("==sum2==", sum2)
}

type T struct {
	name string
}

func compareTest() {
	s := struct{ name string }{"foo"}
	t := T{"foo"}
	fmt.Println(s == t) // true
}

func slicePanic() {
	x := make([]int, 2, 10)
	fmt.Println(x[0:10])
	fmt.Println(x[6])
}

func selectTest() {
	quit := make(chan bool)
	go func() {
		defer fmt.Println("==stop in==")
		for {
			fmt.Println("==select==")
			select {
			case data := <-quit:
				fmt.Println("==data==", data)
			default:
				fmt.Println("==default==")
			}

		}
	}()

	time.Sleep(time.Second * 5)
	quit <- true
	fmt.Println("==stop==")
	time.Sleep(time.Second * 5)

}

func contextTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

type W struct {
}

func (receiver W) eat() {
}

type W2 struct {
}

func (receiver W2) eat() {
}

type Inter1 interface {
	eat()
}
type Inter2 interface {
	eat()
}

func food(dog Inter2) {
	fmt.Println(dog)
}

func interfaceTest() {
	var dog Inter1
	var dog2 Inter2
	fmt.Println(dog == dog2)
	dog = W{}
	dog2 = W2{}
	fmt.Println(dog == dog2)
}

func sizeTest() {
	s := []string{"1", "2", "3", "sadasdasdasdasdasd"}
	s2 := []string{"1"}
	fmt.Println(unsafe.Sizeof(s))  // 24
	fmt.Println(unsafe.Sizeof(s2)) // 24
}

func jsonTest2() {
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市"
   }`)

	data := make(map[string]string)
	json.Unmarshal(jsonData, &data)
	fmt.Println(data)
}

func forSelectTest() {
	ch := make(chan int, 1)
	go func() {
		defer fmt.Println("==defer==")
		a := 0
	EXIT:
		for {
			fmt.Println("==in for==")
			select {
			case data, ok := <-ch:
				a++
				fmt.Println("==data==", data, ok)
				if !ok {
					break EXIT
				}
			}
		}
	}()
	time.Sleep(time.Second * 2)
	ch <- 10
	time.Sleep(time.Second * 2)
	close(ch)
	fmt.Println("==close==")
	time.Sleep(time.Second * 2)
	fmt.Println("==finnal==")
}

type WW struct {
	b int32
	c int64
}

func unsafeTest() {
	var w *WW = new(WW)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age" othertag:"哈哈"`
	score int    `json:"score""`
}

func reflectTest() {
	p := Person{
		"xiaoming",
		20,
		50,
	}
	t := reflect.TypeOf(&p).Elem()
	v := reflect.ValueOf(p)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i), v.Field(i).Interface())
		//fmt.Printf("结构体内第%v个字段 %v 值是%v  对应的json tag是 %v , 还有otherTag？ = %v \n", i+1, t.Field(i).Name, t.Field(i).Type,t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("othertag"))
	}
}

func funcMui(x, y int) (sum int, e error) {
	//sum := x+y
	return sum, nil
}

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func SyncCondTest() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

type Student struct {
	name string
	age  int
}

func atmoicValueTest() {
	var data = atomic.Value{}
	s1 := Student{"xiaoming", 10}
	s2 := Student{"jack", 20}
	data.Store(s1)
	data.Swap(s2)
	fmt.Println(data.Load())
}

func unsafePointerTest() {
	//通过改变
	var a int32
	data := (*int)(unsafe.Pointer(&a))
	fmt.Println(&a)
	fmt.Println(data)

	//	指针运算
	fmt.Println("----------------slice len-----------------------")
	sli := make([]int, 9, 20)
	slen := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sli)) + uintptr(8)))
	fmt.Println(slen, len(sli)) // 9 9
	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sli)) + uintptr(16)))
	fmt.Println(Cap, cap(sli)) // 20 20

	//字符串转数组
	aa := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&aa))
	fmt.Println(ssh)
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v\n", b)
	fmt.Println(&aa)
	fmt.Println(b)
}

func createInt() *int {
	return new(int)
}

func uintptrTest() {
	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y) // 和y一样引用着同一个值
	var p2 = uintptr(unsafe.Pointer(z))

	// 此时，即使z指针值所引用的int值的地址仍旧存储
	// 在p2值中，但是此int值已经不再被使用了，所以垃圾
	// 回收器认为可以回收它所占据的内存块了。另一方面，
	// p0和p1各自所引用的int值仍旧将在下面被使用。

	// uintptr值可以参与算术运算。
	p2 += 2
	p2--
	p2--

	*p0 = 1                         // okay
	*(*int)(p1) = 2                 // okay
	*(*int)(unsafe.Pointer(p2)) = 3 // 危险操作！ 可以通过在后面添加runtime.KeepAlive(z)  就不会出错了
}

func mapTest() {

	fmt.Println("----------------map len-----------------------")
	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18
	lenMap := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(lenMap, len(mp)) // 2 2
}

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func StringToBytes(sa string) []byte {
	return *(*[]byte)(unsafe.Pointer(&sa))
}

func Stringtobyte(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	runtime.KeepAlive(&s)
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func main() {
	str := "abcd"
	a := StringToBytes(str)
	//a := BytesToString([]byte("abcd"))
	//b := []byte(str)

	fmt.Println(a, len(a), cap(a))
	//fmt.Println(b, len(b), cap(b))

	fmt.Println("==string==", (*reflect.StringHeader)(unsafe.Pointer(&str)))
	fmt.Println("==slice==", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	//fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&b)))

	a = append(a, 'a')
	//fmt.Println("==slice[0]==", a[0])
	//a[0] = '1'
	//fmt.Println("==slice==", (*reflect.SliceHeader)(unsafe.Pointer(&a)))

	fmt.Println(a, len(a), cap(a))
	runtime.KeepAlive(&str)

}
