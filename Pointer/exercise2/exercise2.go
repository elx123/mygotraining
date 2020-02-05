package main

import "fmt"

func main(){
	var p *int
	var a int
	p = &a
	fmt.Printf("value address %p the value %d the pointer value %p",p,*p,&p)
}
