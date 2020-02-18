package main

import "fmt"

func main(){
	var stringarray [5]string
	stringarray2 := [5]string{"1","2","3","4","5"}
	stringarray = stringarray2
	for k,_ := range stringarray{
		fmt.Printf("value %s address %p \n",stringarray[k],&stringarray[k])
	}
}
