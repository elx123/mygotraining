package main

import "fmt"

type a interface{
	print()
}

type b interface{
	print()
}

type ab struct{
	name string
}

func(abc ab)print(){
	fmt.Println("123")
}

func main(){
	abcvalue := ab{
		name:"123",
	}
	var A a
	A = abcvalue
	A.print()
	if B,ok := A.(b);ok {
		B.print()
	}
}

