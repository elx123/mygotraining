package main
//https://www.cnblogs.com/my_life/articles/6185890.html
import (
	"encoding/binary"
	"fmt"
)

func main(){
	x := uint32(500)
	buf := make([]byte,4)
	buf2 := make([]byte,4)

	binary.BigEndian.PutUint32(buf,x)

	binary.LittleEndian.PutUint32(buf2,x)
	fmt.Println(buf)
	fmt.Println(buf2)

	x2 := binary.BigEndian.Uint32(buf)
	fmt.Println(x2)

}
