package main

import (
	"fmt"
	"github.com/friendsofgo/errors"
)

type appError struct{
	err error
	message string
	code int
}

func(ae *appError)Error () string{
	return fmt.Sprintf("err %v,message %s,code %d",ae.err,ae.message,ae.code)
}

func(ae *appError)temporary()bool{
	if ae.code == 9{
		return false
	}
	return true
}

func checkFlag(b bool)error{
	if b == false{
		return &appError{errors.New("9"),"31",91}
	}
	return errors.New("flag true")
}

type temporary interface{
	temporary() bool
}

func main(){
	err := checkFlag(false)
	switch e := err.(type){
		case temporary:
			fmt.Println(err)
			if e.temporary(){
				fmt.Println("123456")
			}
	}
}
