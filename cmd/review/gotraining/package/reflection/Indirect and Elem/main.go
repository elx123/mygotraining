package main

import (
	"fmt"
	"reflect"
	"test/test"
)

func main() {
	DifferencePtrAndType()
}

func DifferencePtrAndType() {
	UY := test.Abc{
		Bw: 11,
		Bc: 243,
	}
	YU := &UY
	test := reflect.ValueOf(YU)
	asdf := test.Type().Name()
	iuw := reflect.Indirect(test).Type().Name()
	fmt.Println(asdf)
	fmt.Println(iuw)
}
