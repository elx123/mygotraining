package main

import "fmt"

type speaker interface{
	speak()
	change()
}

type english struct{
	Word string
}

func(e english)change(){
	e.Word = "212"
}

func (e english) speak(){
	fmt.Println(e.Word)
}
//检测4.1.5 Interfaces—Part 2 (Method Sets and Address of Value)  Ultimate Go Programming Second Edition
//视频中的pointer sematic下的value
func main(){
	var spinterface speaker
	en := english{}
	en.Word = "343"
	//ch := chinese{}
	spinterface = &en
	spinterface.speak()
	spinterface.change()
	spinterface.speak()
}
