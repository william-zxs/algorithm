package main

import (
	"errors"
	"fmt"
	"reflect"
)

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
}
