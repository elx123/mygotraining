package main

import (
	"fmt"
	"reflect"
)

type server struct{}

func(sv *server)Test(){
	fmt.Println("123")
}

func TestFunc(f interface{}){
	reflect.ValueOf(f).Call()
	return
}

func main(){
	gg  := &server{}

	f1 :=gg.Test

	TestFunc(f1)
}

