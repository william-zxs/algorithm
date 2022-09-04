package main

import (
	"errors"
	"fmt"
	"reflect"
)

//reflect.Valueof()
//Type() 方法和 Interface() 方法可以打通 interface、Type、Value 三者。
//Type() 方法也可以返回变量的类型信息，与 reflect.TypeOf() 函数等价。
//Interface() 方法可以将 Value 还原成原来的 interface。

//根据 Go 官方关于反射的博客，反射有三大定律：
//
//Reflection goes from interface value to reflection object.
//Reflection goes from reflection object to interface value.
//To modify a reflection object, the value must be settable.
//第一条是最基本的：反射是一种检测存储在 interface 中的类型和值机制。这可以通过 TypeOf 函数和 ValueOf 函数得到。
//
//第二条实际上和第一条是相反的机制，它将 ValueOf 的返回值通过 Interface() 函数反向转变成 interface 变量。
//
//前两条就是说 接口型变量 和 反射类型对象 可以相互转化，反射类型对象实际上就是指的前面说的 reflect.Type 和 reflect.Value。
//
//第三条不太好懂：如果需要操作一个反射变量，那么它必须是可设置的。反射变量可设置的本质是它存储了原变量本身，这样对反射变量的操作，就会反映到原变量本身；反之，如果反射变量不能代表原变量，那么操作了反射变量，不会对原变量产生任何影响，这会给使用者带来疑惑。所以第二种情况在语言层面是不被允许的。
//如果想要操作原变量，反射变量 Value 必须要 hold 住原变量的地址才行。

func reflectDemo() { /**/
	//执行代码会产生 panic，原因是反射变量 v 不能代表 x 本身，为什么？因为调用 reflect.ValueOf(x) 这一行代码的时候，传入的参数在函数内部只是一个拷贝，是值传递，所以 v 代表的只是 x 的一个拷贝，因此对 v 进行操作是被禁止的。
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1) // Error: will panic.
}

//实现通过反射给结构体赋值
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Class
}

type Class struct {
	CName string `json:"cname"`
}

func parseRequest(req interface{}, data map[string]interface{}) error {
	t := reflect.TypeOf(req)
	if t.Kind() != reflect.Pointer {
		return errors.New("need pointer")
	}
	te := t.Elem()

	v := reflect.ValueOf(req)
	ve := v.Elem()
	fmt.Println("==te.NumField()==", te.NumField())
	for i := 0; i < te.NumField(); i++ {
		tef := te.Field(i)
		fmt.Println("==tef==", tef)
		switch tef.Type.Kind() {
		case reflect.String:
			ve.Field(i).SetString((data[tef.Tag.Get("json")]).(string))
		case reflect.Int:
			ve.Field(i).SetInt(int64((data[tef.Tag.Get("json")]).(int)))
		case reflect.Struct:

			fmt.Println("==tef.Type.Kind() ==", tef.Type.Kind())
			for j := 0; j < tef.Type.NumField(); j++ {
				switch tef.Type.Field(j).Type.Kind() {
				case reflect.String:
					ve.Field(i).Field(j).SetString((data[tef.Type.Field(j).Tag.Get("json")]).(string))
				}
			}
		}
	}

	return nil
}

func main() {
	stu := new(Student)
	parseRequest(stu, map[string]interface{}{"name": "jack", "age": 20, "cname": "math"})
	fmt.Println(stu)

	fmt.Printf("")
}
