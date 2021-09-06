package main

import (
	"context"
	"fmt"
	pbGen "github.com/charles-woshicai/golang-advanced-programming/ch4/pubsub/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"testing"
	"time"
)

func TestRunPubsubServer(t *testing.T) {
	lis := runPubsubServer()
	defer lis.Close()

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pbGen.NewPubsubServiceClient(conn)

	// subscribe
	stream, err := client.Subscribe(
		context.Background(), &pbGen.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	// publish
	_, err = client.Publish(
		context.Background(), &pbGen.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &pbGen.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
	reply, err := stream.Recv()
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	fmt.Println(reply.GetValue())
}
