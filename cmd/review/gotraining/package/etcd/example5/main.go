package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
)

func main() {
	var cli *clientv3.Client
	var err error

	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err = clientv3.New(config)
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	cli.KV = namespace.NewKV(cli.KV, "test")

	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	gresp, err := cli.Get(context.TODO(), "foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gresp)
}
