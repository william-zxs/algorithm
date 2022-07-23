package common

//翻转含有中⽂、数字、英⽂字⺟的字符串
func ReverseStr(str string) string {
	//注意点是又中文字符，所以要先转成[]rune

	//"你好abc啊哈哈213"
	//"你a2"
	//"你a22"
	s := []rune(str)
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return string(s)
}
