package test

import "fmt"

//as
type Abc struct {
	Bw int
	Bc float64
}

func (*Abc) SayHello() {
	fmt.Println("123")
}
