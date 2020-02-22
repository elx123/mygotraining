package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)


var sendmessage chan []byte
var receiveMessage chan []byte

func init() {
	sendmessage = make(chan []byte)

	receiveMessage = make(chan []byte)

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

//这里我们提出一个疑问socket支持并行读与写吗
//https://bbs.csdn.net/topics/60501063
//可以一边读一边写（全双工）
func main(){
	conn,err := net.Dial("tcp","127.0.0.1:12345")
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go Woker(conn,&wg)
	go terminalRead()
	go ReadSocket(conn)
	wg.Wait()
}

func Woker(cn net.Conn,wg *sync.WaitGroup){
	defer cn.Close()
	for {
		select {
			case sendbuf  := <-sendmessage:
				{
					log.Println("begin to send ",sendbuf)
					n,err :=cn.Write(sendbuf)
					if err!= nil{
						fmt.Print(err)
						break
					}
					log.Print(n)
				}
			case receivebuf := <-receiveMessage:
				{
					log.Print(string(receivebuf))
				}
		}
	}
	wg.Done()
}

func terminalRead() {
	buf := make([]byte,1)
	result := make([]byte,0,100)
	for {
		_,err := io.ReadFull(os.Stdin,buf)
		if err != nil{
			break
		}
		if buf[0] == '\n' {
			result = append(result,buf[0])
			log.Println(result)
			sendmessage <- result
			result = result[:0]
			continue
		}
		result = append(result,buf[0])
	}
}

func ReadSocket(cn net.Conn){
	buf := make([]byte,1)
	socketBuf := make([]byte,0,100)
	for {
		_,err := io.ReadFull(cn,buf)
		if err != nil {
			log.Println(err)
			break
		}
		if buf[0] == '\n'{
			socketBuf = append(socketBuf,buf[0])
			receiveMessage <- socketBuf
			socketBuf = socketBuf[:0]
			continue
		}
		socketBuf = append(socketBuf,buf[0])
	}
	cn.Close()
}