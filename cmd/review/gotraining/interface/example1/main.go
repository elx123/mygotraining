package main

import "fmt"

type TestABC interface{
	Test()
}

type TestStrcut struct{
	TestABC
	key,value interface{}
}

func(t *TestStrcut)Test(){
	fmt.Println(t.key,t.value)
}

func main(){
	root := new(TestStrcut)
	second := TestStrcut{root,123,456}
	second.Test()
	third := TestStrcut{&second,798,765}
	third.Test()
}

