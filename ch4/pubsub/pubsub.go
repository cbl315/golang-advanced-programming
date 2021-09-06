package main

import (
	"context"
	"fmt"
	pbGen "github.com/charles-woshicai/golang-advanced-programming/ch4/pubsub/pb"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
	pbGen.PubsubServiceServer
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(
	ctx context.Context, arg *pbGen.String,
) (*pbGen.String, error) {
	fmt.Printf("publish topic: %s\n", arg.GetValue())
	p.pub.Publish(arg.GetValue())
	return &pbGen.String{}, nil
}

func (p *PubsubService) Subscribe(arg *pbGen.String, stream pbGen.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		fmt.Println("receive topic", v)
		if err := stream.Send(&pbGen.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func runPubsubServer() net.Listener {
	grpcServer := grpc.NewServer()
	pbGen.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	go grpcServer.Serve(lis)
	return lis
}

func main() {

}
