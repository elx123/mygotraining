package main

import "fmt"

func main(){
	testmap := make(map[int]string)
	testmap[1] = "1"
	testmap[2] = "2"
	testmap[3] = "3"
	testmap[4] = "4"
	testmap[5] = "5"
	for k,v:= range testmap{
		fmt.Println(k,v)
	}
}
