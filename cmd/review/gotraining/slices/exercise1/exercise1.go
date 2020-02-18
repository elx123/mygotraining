package main

import "fmt"

func main(){
	var slice1 []int
	for i:=0;i<10;i++{
		slice1 = append(slice1,i)
	}
	for k,v := range slice1{
		fmt.Printf("key %d,value %d \n",k,v)
	}
}
