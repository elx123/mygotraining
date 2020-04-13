package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var test = strings.NewReplacer("a", "A", "b", "B")
	s := test.Replace("acpbsf")
	fmt.Println(s) //AcpBsf

	w, size := utf8.DecodeRuneInString("国家3r")
	fmt.Println(w, size)

	fmt.Printf("%q", w)
}
