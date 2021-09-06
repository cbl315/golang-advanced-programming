package main

import (
	"context"
	pbGen "github.com/charles-woshicai/golang-advanced-programming/ch4/gRPC/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloServiceImpl struct {
	pbGen.HelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pbGen.String) (*pbGen.String, error) {
	reply := &pbGen.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream pbGen.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pbGen.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func runGRPC() net.Listener {
	grpcServer := grpc.NewServer()
	pbGen.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	go grpcServer.Serve(lis)
	return lis
}

func main() {
	runGRPC()
}
