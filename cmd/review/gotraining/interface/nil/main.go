package main

import "fmt"

func main() {
	var f interface{}

	if f == nil {
		fmt.Println("123")
	} else {
		fmt.Println(f)
	}
}
