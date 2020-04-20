package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func test(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/test", test)

	server.ListenAndServe()
}
