package main

import "fmt"

func main() {

	fmt.Println(test("adsfdsf"))

}

func test(s string) int {
	n := len(s)
	var len int
	result := make(map[byte]int)
	tempstring := []byte(s)
	var i, j int
	for i < n && j < n {
		if _, ok := result[tempstring[j]]; !ok {
			result[tempstring[j]] = j
			j++
			if j-i > len {
				len = j - i
			}
		} else {
			delete(result, tempstring[i])
			i++
		}
	}
	return len
}
