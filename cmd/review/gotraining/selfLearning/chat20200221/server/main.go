package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

var players []net.Conn
var enterchannel chan string
var messagechannel chan []byte
var leavechannel chan string

func init() {
	players = make([]net.Conn,0,100)

	enterchannel = make(chan string,100)
	// Change the output device from the default
	// stderr to stdout.
	log.SetOutput(os.Stdout)

	// Set the prefix string for each log line.
	log.SetPrefix("TRACE: ")

	// Set the extra log info.
	setFlags()
}

func setFlags() {
	/*
	   Ldate			// the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//对于进入房间的人一旦有人进入房间,就进行全服通告,一旦有人离开房间,也一样,同时还会全服通告语言


func main(){
	ln,err := net.Listen("tcp","0.0.0.0:12345")
	if err != nil{
		log.Fatal(err)
	}
	go handleEnterUser()
	for{
		conn,err := ln.Accept()
		if err != nil{
			conn.Close()
			continue
		}
		enterchannel <- conn.RemoteAddr().String()
		players = append(players,conn)
		log.Print(conn.RemoteAddr().String()+" entered")
		go handle(conn)
	}
}

func handle(cn net.Conn){
	defer cn.Close()
	var sendbuf bufio.ReadWriter
	buf := make([]byte,0,1)
	for {
		_,err := cn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Print(buf)
		if buf[0] == '\n'{
			var buf []byte
			sendbuf.Write(buf)

			messagechannel <- buf
		}
		sendbuf.Write(buf)
	}
	leavechannel <- cn.RemoteAddr().String()
}

func handleEnterUser() {
	for {
		select {
		case user := <-enterchannel:
			for i := range players {
				io.WriteString(players[i], user)
			}
		case message := <-messagechannel:
			{
				for i := range players {
					players[i].Write(message)
				}
			}
		case user := <-leavechannel:
			{
				for i := range players {
					io.WriteString(players[i], user)
				}
			}
		}
	}
}
