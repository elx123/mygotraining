package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	r := strings.NewReader("/directory/myFollow")
	resp,err := http.Post("http://www.douyu.com", "text/plain", r)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)
}
