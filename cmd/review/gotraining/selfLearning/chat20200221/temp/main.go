package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := make([]byte,2)
	//buf[0] = '\n'
	//buf[1] = '\n'
	stringbuf := string(buf)
	fmt.Println(stringbuf)
	bbuf := bytes.Buffer{}
	bbuf.Write([]byte("阿打算考几分开了都是你发的时刻"))
	r,z,err := bbuf.ReadRune()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%+q%v",r,z)
	err = bbuf.UnreadRune()
	if err != nil{
		fmt.Println(err)
	}
	r,z,err = bbuf.ReadRune()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%+q%v",r,z)
}

