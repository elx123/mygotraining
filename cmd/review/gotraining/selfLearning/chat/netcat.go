// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
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
	"os"
	"strings"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {

		input := bufio.NewReaderSize(conn, 65536)
		_, err_loop := input.Peek(4)
		for ; err_loop != io.EOF; _, err_loop = input.Peek(4) {
			if err_loop == bufio.ErrBufferFull {
				continue
			} else if err_loop != nil {
				break
			} else {
				if input.Buffered() >= 4 {
					var len_heard int32
					heard_len, ok := input.Peek(4)
					if ok != nil {
						log.Fatal("error1")
					}
					buf := bytes.NewReader(heard_len)
					err := binary.Read(buf, binary.LittleEndian, &len_heard)
					if err != nil {
						log.Fatal("error2")
					}
					input.Discard(4)
					if len_heard > 65536 || len_heard < 0 {
						log.Printf("Invalid length :%d", len_heard)
						conn.Close()
						break
					} else if int32(input.Buffered()) >= (len_heard) {
						message_byte, ok := input.Peek(int(len_heard))
						if ok != nil {
							log.Fatal("error3")
						}
						//messages <- who + ": " + string(message_byte)
						fmt.Println(string(message_byte))
						input.Discard(int(len_heard))
					} else {
						break
					}
				}
			}
		}
		//io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	reader := bufio.NewReader(os.Stdin)
	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			break
		}
		result = strings.Replace(result, "\n", "", -1)
		result_byte, err2 := codec.Encode(result)
		if err2 != nil {
			log.Print(err2)
			break
		}
		conn.Write(result_byte)
	}
	conn.Close()
	<-done
}
