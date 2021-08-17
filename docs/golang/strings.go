package golang

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
ExampleString:

golang string存储转换过程,字符集配合编码模式

string >  unicode+编码模式（utf-8） > binary

定长编码:
	浪费内存，字符位数越长，浪费越严重
变长编码:
	代表就是utf8编码,是golang默认的编码模式
	区间：       字节数：     编码模板：
	[0,127]     占用1字节    0???????
	[128,2037]  占用2字节    110????? 10??????
	[2048,65535]占用3字节    110????? 10?????? 10??????
	example: g 103 01100111

golang认为string是不会被修改，会将string分配到只读代码段
string本质：

// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.

type StringHeader struct {
	Data uintptr  指向数据的内存地址
	Len  int      string的长度
}


注意事项：
	string 在转换 []byte 时会发生内存拷贝  []byte(string)

	string = StringHeader
	[]byte = SliceHeader
	通过修改内存指向的地址，可以实现不拷贝进行stirng和[]byte的类型转换

func str2bytes(s string) []byte {
	r := (*reflect.StringHeader)(unsafe.Pointer(&s))
	res := reflect.SliceHeader{
		Data: r.Data,
		Len:  r.Len,
		Cap:  r.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&res))
}

func bytes2str(b []byte) string {
	r := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	str := reflect.StringHeader{
		Data: r.Data,
		Len:  r.Len,
	}
	return *(*string)(unsafe.Pointer(&str))
}
*/
func ExampleString() {
	str := "gggg"

	res := str2bytes(str)

	r := bytes2str(res)
	fmt.Println(str, r)
	fmt.Println(str == r)

}


func str2bytes(s string) []byte {
	r := (*reflect.StringHeader)(unsafe.Pointer(&s))
	res := reflect.SliceHeader{
		Data: r.Data,
		Len:  r.Len,
		Cap:  r.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&res))
}

func bytes2str(b []byte) string {
	r := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	str := reflect.StringHeader{
		Data: r.Data,
		Len:  r.Len,
	}
	return *(*string)(unsafe.Pointer(&str))
}

