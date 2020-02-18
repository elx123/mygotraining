//coder with chat server

package main

import (
	"bufio"
	"bytes"
	"codec"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	input := bufio.NewReaderSize(conn, 65536)
	_, err_loop := input.Peek(4)
	var len_heard int32

	for ; err_loop != io.EOF; input.Peek(int(len_heard)) {
		log.Print("hhhhhhhh")
		if err_loop == bufio.ErrBufferFull {
			log.Println("test3")
			continue
		} else if err_loop != nil {
			log.Println("test4")
			break
		} else {
			if input.Buffered() >= 4 {
				heard_len, ok := input.Peek(4)
				if ok != nil {
					log.Fatal("error")
				}
				buf := bytes.NewReader(heard_len)
				err := binary.Read(buf, binary.LittleEndian, &len_heard)
				if err != nil {
					log.Fatal("error")
				}
				if len_heard > 65536 || len_heard < 0 {
					log.Printf("Invalid length :%d", len_heard)
					conn.Close()
					break
				} else if int32(input.Buffered()-4) >= (len_heard) {
					input.Discard(4)
					message_byte, ok := input.Peek(int(len_heard))
					if ok != nil {
						log.Fatal("error")
					}
					messages <- who + ": " + string(message_byte)
					fmt.Println(string(message_byte))
					input.Discard(int(len_heard))
					//time.Sleep(time.Millisecond*100)
				} else {
					continue
				}
			}
		}
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		result_byte, err2 := codec.Encode(msg)
		if err2 != nil {
			log.Fatalf("encode error:", err2)
		}
		conn.Write(result_byte)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
