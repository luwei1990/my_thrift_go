package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"helloworld"
	"context"
)

func main(){

	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket("localhost:9090")
	if err != nil{
		fmt.Println("NewTsocket failed ", err)
		return
	}

	transport, err = thrift.NewTBufferedTransportFactory(
		8192).GetTransport(transport)

	if err != nil {
		fmt.Errorf("NewTransport failed. err: [%v]\n", err)
		return
	}


	defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Errorf("Transport.Open failed. err: [%v]\n", err)
		return
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := helloworld.NewHelloWorldClient(thrift.NewTStandardClient(iprot, oprot))

	res, err := client.Ping(context.Background())
	if err != nil{
		fmt.Println("ping error")
		return
	}

	fmt.Println("message from server ", res)


	res, err = client.Say(context.Background(), "luwei")
	if err != nil{
		fmt.Println("say hello failed")
		return
	}
	fmt.Println("message from server ", res)

}