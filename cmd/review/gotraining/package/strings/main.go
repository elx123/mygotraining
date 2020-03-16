package main

import (
	"fmt"
	"strings"
)

func main() {
	var test = strings.NewReplacer("a", "A", "b", "B")
	s := test.Replace("acpbsf")
	fmt.Println(s) //AcpBsf
}
