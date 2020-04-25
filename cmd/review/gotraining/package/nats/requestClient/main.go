package main

import (
	"flag"
	"log"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/pborman/uuid"
)

const (
	//url   = "nats://192.168.3.125:4222"
	url = "nats://192.168.99.100:30009"
)

var (
	nc         *nats.Conn
	encodeConn *nats.EncodedConn
	err        error
)

func init() {
	if nc, err = nats.Connect(url); checkErr(err, func() {

	}) {
		//
		if encodeConn, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER); checkErr(err, func() {

		}) {

		}
	}
}

func main() {
	var (
		subj = flag.String("subj", "yasenagat", "subject name")
	)
	flag.Parse()
	log.Println(*subj)
	startClient(*subj)

	time.Sleep(time.Second)
}

//send message to server
func startClient(subj string) {
	for i := 0; i < 3; i++ {
		id := uuid.New()
		log.Println(id)
		if msg, err := nc.Request(subj, []byte(id+" hello"), time.Second); checkErr(err, func() {
			// handle err
		}) {
			log.Println(string(msg.Data))
		}
	}
}

func checkErr(err error, errFun func()) bool {
	if err != nil {
		log.Println(err)
		errFun()
		return false
	}
	return true
}
