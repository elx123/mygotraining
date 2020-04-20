package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func test(w http.ResponseWriter, req *http.Request) {
	err := db.Ping()
	if err != nil {
		w.Write([]byte("Success"))
	} else {
		w.Write([]byte("Fail"))
	}
	return
}

func main() {

	db, err := sql.Open("mysql", "root:123456sdfwreS@tcp(mysqldb:3306)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/test", test)

	server.ListenAndServe()
}
