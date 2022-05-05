package main

import (
	"testing"
	"unsafe"
)

func chStr2Sl1(str string) []byte {
	return []byte(str)
}

func chStr2Sl2(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

var str = "转换后 [ ]byte 底层数ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否ring 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否" +
	"组与原 string 内部指针并不相同，以此可确定数据被复制。那么，如不修改数据，仅转换类型，是否可避开复制，从而提升性能"

func BenchmarkChStr2Sl1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chStr2Sl1(str)
	}
}

func BenchmarkChchStr2Sl2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chStr2Sl2(str)
	}
}
