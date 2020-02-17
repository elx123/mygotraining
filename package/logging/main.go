package main

import (
	"fmt"
	"log"
)

func init(){
	log.SetPrefix("jiajun:")
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}

func main(){
	log.Println("12345")
	fmt.Println(123)
}
