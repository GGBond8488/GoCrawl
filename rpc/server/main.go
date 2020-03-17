package main

import (
	rpcdemo "WebPas/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn ,err := listener.Accept()
		if err != nil{
			log.Printf("accpet error: %v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
