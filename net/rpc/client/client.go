package main

import (
	localrpc "basic/net/rpc"
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ localrpc.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("rpc dial error:", err)
	}

	return &HelloServiceClient{c}, nil
}

func (c *HelloServiceClient) Hello(request string, reply *string) error {
	return c.Call(localrpc.HelloServiceName+".Hello", request, reply)
}

func main1() {
	client, err := DialHelloService("tcp", "localhost:1234")
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
