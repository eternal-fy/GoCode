package main

import (
	"learnCode/net/rpc/rpc_object"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {

	calc := new(rpc_object.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
