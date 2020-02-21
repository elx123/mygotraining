package main

import (
	"bytes"
	"os"
)

func main(){
	var bbuf bytes.Buffer
	for {
		buf := make([]byte,1)
		os.Stdin.Read(buf)
		if buf[0] == '\n'{
			bbuf.Write(buf)
			break
		}
		bbuf.Write(buf)
	}
	bbuf.WriteTo(os.Stdout)
}
