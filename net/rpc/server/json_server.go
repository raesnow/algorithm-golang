package main

import (
	localrpc "basic/net/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceJson struct {
}

func (s *HelloServiceJson) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func RegisterHelloServiceJson(svc localrpc.HelloServiceInterface) error {
	return rpc.RegisterName(localrpc.HelloServiceName, svc)
}

func main2() {
	RegisterHelloService(new(HelloServiceJson))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
