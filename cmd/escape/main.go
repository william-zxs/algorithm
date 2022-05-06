package main

func newString() *string {
	s := new(string)
	*s = "wohu"
	return s
}

//func Escapetest() {
//var s string
//s := new(string)
//_ = newString(s)
//fmt.Println(s)
//}

func main() {
	newString()
	//Escapetest()
}
