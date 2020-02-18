package main

import "fmt"

func main(){
	names :=[]string{"12","22","33","44","55"}

	// Take a slice of index 1 and 2.
	slice := names[1:3]

	// Display the value of the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
