package main

import (
	localrpc "basic/net/rpc"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func RegisterHelloService(svc localrpc.HelloServiceInterface) error {
	return rpc.RegisterName(localrpc.HelloServiceName, svc)
}

func main1() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
