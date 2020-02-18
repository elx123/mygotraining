package main

import (
	"fmt"
	"context"
	"time"
)

func main(){

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx2,cancel2 := context.WithTimeout(ctx,5*time.Second)
	defer cancel2()

	//ctx3,cancel3 := context.WithDeadline(ctx2,time.Now())
	//defer cancel3()

	ctx3:= context.WithValue(ctx2,123,321)

	<- ctx3.Done()
	fmt.Println(ctx3.Value(123))
}
