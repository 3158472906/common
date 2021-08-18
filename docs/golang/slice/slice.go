package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	ExampleSlice

slice的本质:
	// It cannot be used safely or portably and its representation may
	// change in a later release.
	// Moreover, the Data field is not sufficient to guarantee the data
	// it references will not be garbage collected, so programs must keep
	// a separate, correctly typed pointer to the underlying data.
	type SliceHeader struct {
		Data uintptr  指向内存地址
		Len  int      元素个数
		Cap  int	  元素容量，与开辟的内存地址
	}

slice注意事项：
	1. 通过下标访问数组时注意角标越界错误
	2. 使用make([]type,len,cap)的方式初始化slice
	3. 初始化时尽可能的确定slice的cap，避免slice扩容
	4. len([]type) 来获取slice的元素个数
	5. cap([]type) 来获取slice的容量
	6. []type = append([]type,args...) 来给slice追加数据
	7. copy(dest,src) 来进行slice的拷贝

slice的扩容规则：
	第一步 预估容量
		oldCap表示生命slice时的容量
		currentCap = cap([]type)
		newCap = oldCap * 2
		if newCap < currentCap
				newCap = currentCap
		else if oldLen < 1024
				newCap = oldCap*2
		else if oldLen > 1024
				newCap = oldCap*1.25

	第二部 计算当前slice所需要的内存大小

	第三步 内存对齐,内存匹配
		申请内存时向golang的内存管理模块申请内存
		golang的内存管理模块会预先分配好很多内存块，都是2的倍数
		然后返回足够大而且最接近slice的内存大小的规格，并计算出新的cap
*/
func ExampleSlice() {
	data := []string{"1", "2", "3"}

	r := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	fmt.Println(r.Data)
	fmt.Println(r.Len)
	fmt.Println(r.Cap)

	// 结构体内存对齐
	s := St{}
	fmt.Println(unsafe.Sizeof(s))
}

type St struct {
	A int8  // 1
	D int16 // 2
	C int32 // 4
	B int64 // 8
}
