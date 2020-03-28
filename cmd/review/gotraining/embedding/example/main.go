//这个例子讲了2个问题
//parent的cancel会影响child context  第一个
//only for request-scoped data   第二个
package main

import "fmt"

type ABC interface {
	testabc()
	testcba()
}

type TestAbcS struct {
	ABC
	A interface{}
	B interface{}
}

func (t TestAbcS) testabc() {
	fmt.Println("123")
}

func main() {
	first := TestAbcS{
		A: 1,
		B: 1,
	}

	var firstInterface ABC

	firstInterface = first

	second := TestAbcS{
		ABC: firstInterface,
		A:   1,
		B:   1,
	}
	fmt.Println(second.A)
}
