package main

import "reflect"

//go中会panic的情况

func func1() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1) // Error: will panic.
}

//向关闭的channel写
//close channel两次

func main() {

}
