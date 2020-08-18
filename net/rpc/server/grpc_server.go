package main

import (
	"basic/net/rpc"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloServiceImpl struct {
}

func (s *HelloServiceImpl) Hello(ctx context.Context, args *rpc.String) (*rpc.String, error) {
	reply := &rpc.String{
		Value: "Hello, " + args.GetValue(),
	}
	return reply, nil
}

func (s *HelloServiceImpl) Channel(stream rpc.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &rpc.String{Value: "hello, " + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return nil
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	rpc.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(listener)
}
