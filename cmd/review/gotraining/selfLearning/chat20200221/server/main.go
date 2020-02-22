package main

import (
	"io"
	"log"
	"net"
	"os"
)

var players map[net.Conn]interface{}
var enterchannel chan net.Conn
var messagechannel chan []byte
var leavechannel chan net.Conn

func init() {
	players = make(map[net.Conn]interface{})

	enterchannel = make(chan net.Conn,100)
	messagechannel = make(chan []byte)
	leavechannel = make(chan net.Conn,100)
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
		enterchannel <- conn
		go handle(conn)
	}
}

func handle(cn net.Conn){
	defer cn.Close()
	receivebuf := make([]byte,0,100)
	buf := make([]byte,1)
	for {
		_,err := io.ReadFull(cn,buf)
		if err != nil {
			log.Println(err)
			break
		}
		if buf[0] == '\n'{
			receivebuf = append(receivebuf,buf[0])
			messagechannel <- receivebuf
			receivebuf = receivebuf[:0]
			continue
		}
		receivebuf = append(receivebuf,buf[0])
	}
	leavechannel <- cn
}

func handleEnterUser() {
	for {
		select {
		case user := <-enterchannel:
			log.Print(user.RemoteAddr().String()+" entered")
			log.Print("player num ",len(players))
			for i := range players {
				io.WriteString(i, user.RemoteAddr().String()+" entered\n")
			}
			players[user] = 1
		case message := <-messagechannel:
			log.Print("player num ",len(players))
			{
				for i := range players {
					i.Write(message)
				}
			}
		case user := <-leavechannel:
				log.Print(user.RemoteAddr().String()+" leave")
				delete(players,user)
				for i := range players {
					io.WriteString(i, user.RemoteAddr().String()+" leave\n")
				}

		}
	}
}
