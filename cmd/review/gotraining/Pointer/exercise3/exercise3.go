package main

import "fmt"

type usertype struct{
	a int
	b int
	c string
}

func main(){
	ut := usertype{
		a: 0,
		b: 0,
		c: "123",
	}
	fmt.Printf("%v",ut)
	test(&ut)
	fmt.Printf("%v",ut)
}

func test(pusertype *usertype){
	pusertype.a = 1
}