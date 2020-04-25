package main

import (
	"flag"
	"log"

	"github.com/nats-io/go-nats"
)

const (
	//url   = "nats://192.168.3.125:4222"
	url = "nats://192.168.99.100:30009"
)

var (
	nc *nats.Conn

	encodeConn *nats.EncodedConn
	err        error
)

func init() {
	if nc, err = nats.Connect(url); checkErr(err) {
		//
		if encodeConn, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER); checkErr(err) {
		}
	}
}

func main() {
	var (
		servername = flag.String("servername", "Y", "name for server")
		queueGroup = flag.String("group", "", "group name for Subscribe")
		subj       = flag.String("subj", "yasenagat", "subject name")
	)
	flag.Parse()

	mode := "queue"
	if *queueGroup == "" {
		mode = "pub/sub"
	}
	log.Printf("Server[%v] Subscribe Subject[%v] in [%v]Mode", *servername, *subj, mode)

	startService(*subj, *servername+" worker1", *queueGroup)
	startService(*subj, *servername+" worker2", *queueGroup)
	startService(*subj, *servername+" worker3", *queueGroup)

	nc.Flush()
	select {}
}

//receive message
func startService(subj, name, queue string) {
	go async(nc, subj, name, queue)
}

func async(nc *nats.Conn, subj, name, queue string) {
	replyMsg := name + " Received a msg"
	if queue == "" {
		nc.Subscribe(subj, func(msg *nats.Msg) {
			nc.Publish(msg.Reply, []byte(replyMsg))
			log.Println(name, "Received a message From Async : ", string(msg.Data))
		})
	} else {
		nc.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
			nc.Publish(msg.Reply, []byte(replyMsg))
			log.Println(name, "Received a message From Async : ", string(msg.Data))
		})
	}

}

func checkErr(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
