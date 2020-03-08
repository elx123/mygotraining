package main

import (
	"context"
	"fmt"
	"log"
	"mygotraining/cmd/grpcTest/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "alx",
			LastName:  "cjj",
		},
	}

	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(res.Result)
}
