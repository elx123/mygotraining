package main

import (
	"fmt"
)

func main(){
// 首先要理解gotraining中slice课程
	buf := make([]byte,10)
	fmt.Println(buf,len(buf),cap(buf))
	fmt.Printf("%p\n",&buf[0])
	fmt.Printf("%p\n",&buf)
	fmt.Printf("%p\n",buf)
	//slice作为引用类型,他的地址和slice中元素的地址是不同的

	//如何去删除slice中元素
	buf = buf[:0]
	fmt.Println(buf,len(buf),cap(buf))

	//如何给slice扩容
	buf = buf[:5]
	fmt.Println(buf,len(buf),cap(buf))

	//可以直接将string copy进入slice
	n := copy(buf,"国")
	fmt.Println(buf,n,len(buf),cap(buf))
}
