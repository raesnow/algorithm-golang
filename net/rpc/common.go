package rpc

const HelloServiceName = "pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}
