package main

import "fmt"

func main() {
	buf := make([]byte,2)
	//buf[0] = '\n'
	//buf[1] = '\n'
	stringbuf := string(buf)
	fmt.Println(stringbuf)
}

