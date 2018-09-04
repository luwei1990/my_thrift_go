package main

import "fmt"
import "git.apache.org/thrift.git/lib/go/thrift"
import (
	"helloworld"
	"context"
)

type HelloWorld struct {
	
}

func (h *HelloWorld)Ping(ctx context.Context) (string, error){
	fmt.Println("Pong")
	return "Pong", nil
}

func (h *HelloWorld)Say(ctx context.Context, msg string) (string, error)  {
	fmt.Println("get string", msg)
	return "Send: " + msg, nil
}

func main()  {
	transport, err := thrift.NewTServerSocket("localhost:9090")
	if err != nil{
		panic(err)
	}

	processor := helloworld.NewHelloWorldProcessor(&HelloWorld{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTBinaryProtocolFactoryDefault(),
	)
	if err := server.Serve(); err != nil{
		panic(err)
	}

}

