package main

import (
	bufio2 "bufio"
	"bytes"
	"fmt"
)

func main(){
	buf := bytes.Buffer{}
	bufio := bufio2.Writer{}
	temp := "国"
	buf.Write([]byte(temp))
	r,n,err := buf.ReadRune()
	fmt.Println(r,n,err)
	fmt.Printf("%+q %d %v",r,n,err)
}
