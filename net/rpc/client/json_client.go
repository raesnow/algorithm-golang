package main

import (
	localrpc "basic/net/rpc"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceClientJson struct {
	*rpc.Client
}

var _ localrpc.HelloServiceInterface = (*HelloServiceClientJson)(nil)

func DialHelloServiceJson(network, address string) (*HelloServiceClientJson, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		log.Fatal("rpc dial error:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))

	return &HelloServiceClientJson{client}, nil
}

func (c *HelloServiceClientJson) Hello(request string, reply *string) error {
	return c.Call(localrpc.HelloServiceName+".Hello", request, reply)
}

func main2() {
	client, err := DialHelloServiceJson("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("rpc dial error:", err)
	}

	var reply string
	err = client.Hello("buddy", &reply)
	if err != nil {
		log.Fatal("client call error:", err)
	}

	fmt.Println("Reply:", reply)
}
