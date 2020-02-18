package Exercise2

import (
	"fmt"
	"math/rand"
)

func Exercise2() {
	ch := make(chan int,100)
	for i:=0;i<100;i++{
		go func(){
			n := rand.Intn(1000)
			ch <- n
		}()
	}
	num := 0
	for kk:=0;kk<100;kk++{
		fmt.Println(num)
		nn := <- ch
		fmt.Println(nn)
		num++
	}
}