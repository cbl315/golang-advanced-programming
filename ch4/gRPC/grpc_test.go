package main

import (
	"context"
	"fmt"
	pbGen "github.com/charles-woshicai/golang-advanced-programming/ch4/gRPC/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"testing"
	"time"
)

func Test_runGRPC(t *testing.T) {
	lis := runGRPC()
	defer lis.Close()
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Hello func
	client := pbGen.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &pbGen.String{Value: "hello gRPC"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
	//Channel func
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	stopCh := make(chan struct{})
	go func(stopCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				return
			default:
				if err := stream.Send(&pbGen.String{Value: "hi"}); err != nil {
					log.Fatal(err)
				}
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	for i := 0; i < 5; i++ {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
	stopCh <- struct{}{}
}
