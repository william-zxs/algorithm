package main

import (
	"fmt"
	"reflect"
)

//Golang Comparable Types

//Boolean values
//Integer values
//Floating point values
//Complex values
//String values
//Pointer values
//Channel values
//Interface values
//Struct values are comparable if all their fields are comparable
//Array values are comparable if values of the array element type are comparable
//A value x of non-interface type X and a value t of interface type T are comparable when values of type X are comparable and X implements T

// ==比较

//两个接口比较，比较的是动态类型和动态值，这两个都一样才相等

//func, slice, map 不可以用==比较，只能和nil比较
//对于map和slice 可以用 reflect.DeepEqual 比较

func compareFunc() {
	//直接用等号比较会编译失败
	//a := func() {}
	//b := func() {}
	//fmt.Println(a == b)

	//用DeepEqual  结果永远是false
	//a := func() {}
	//b := func() {}
	//fmt.Println(reflect.DeepEqual(a, b))
}

type S struct {
	//A func()
}

func compareStruct() {
	s1 := S{}
	//s2 := S{}
	v := reflect.TypeOf(s1).Comparable()
	fmt.Println(v)
}
func main() {
	//compareFunc()
	compareStruct()
}
