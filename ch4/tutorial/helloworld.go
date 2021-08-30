package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

type CloseFunc func()

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func Run() CloseFunc {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		rpc.ServeConn(conn)
	}()
	return func() {
		_ = listener.Close()
	}
}

func RunJsonRpc() CloseFunc {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}()
	return func() {
		_ = listener.Close()
	}
}

func main() {
	Run()
}
