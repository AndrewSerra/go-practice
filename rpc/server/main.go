package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc_example/command"
)

func main() {
	server := new(command.TaskServer)

	rpc.Register(server)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error: ", err)
	}

	http.Serve(listener, nil)
}
