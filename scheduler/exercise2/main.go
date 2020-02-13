package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	a := func(){
		for i:= 100;i>0;i--{
			fmt.Println("a:",i)
		}
		wg.Done()
	}
	b := func(){
		for i:=0;i<100;i++{
			fmt.Println("b:",i)
		}
		wg.Done()
	}
	go a()
	go b()
	wg.Wait()
}
