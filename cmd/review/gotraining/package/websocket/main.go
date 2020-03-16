package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func websocket(w http.ResponseWriter, r *http.Request)){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	
	ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
	}
	ws,err := upgrader
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/websocket", headers)
	server.ListenAndServe()
}