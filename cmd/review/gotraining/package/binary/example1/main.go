package example1
//结合例子判断大小端理解大小端
import (
	"fmt"
	"unsafe" //go语言的sizeof
)

func main() {
	s := int16(0x1234)
	b := int8(s)
	fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
	if 0x34 == b {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}
}
