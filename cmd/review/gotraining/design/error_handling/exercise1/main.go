package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidValue = errors.New("ErrInvalidValue")
	ErrAmountTooLarge = errors.New("ErrAmountTooLarge")
)

func checkAmount(a float64)error{
	if a == 0{
		return ErrInvalidValue
	}
	if a == 1000 {
		return ErrAmountTooLarge
	}
	return nil
}

func main(){
	err := checkAmount(0)
	switch err {
		case ErrInvalidValue :{
			fmt.Println("is ErrInvalidValue")
		}
		case ErrAmountTooLarge:{
			fmt.Println("is ErrAmountTooLarge")
		}
		default:
			fmt.Println("unknow err")
	}
	return
}
