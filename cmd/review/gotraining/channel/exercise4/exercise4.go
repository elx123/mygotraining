package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init(){
	rand.Seed(time.Now().Unix())
}

func main(){
	num := runtime.NumCPU()
	ch := make(chan int)
	shutdown := make(chan struct{})
	result := make([]int,0,100)
	var wg sync.WaitGroup
	wg.Add(num)
	for i:=0;i<num;i++{
		go func(id int){
			defer wg.Done()
			for{

				randnum := rand.Intn(1000)
				select{
					case ch <- randnum :
						fmt.Printf("Gocorunine %d generate num %d \n",id,randnum)
					case <- shutdown :
						fmt.Printf("Gocorunine %d shutdown \n",id)
						return
				}
			}
		}(i)
	}

	for n := range ch {
		if n%2 == 0{
			continue
		}
		result = append(result,n)
		fmt.Printf("len %d \n",len(result))
		if len(result) == 100{
			break
		}
	}
	close(shutdown)
	wg.Wait()
}
