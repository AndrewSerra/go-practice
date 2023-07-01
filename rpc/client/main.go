package main

import (
	"log"
	"net/rpc"
	"rpc_example/command"
)

func main() {

	var reply bool
	data := command.TaskData{
		UserId:      "asdf",
		ProjectName: "Welcome",
	}

	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	err = client.Call("TaskServer.Remove", data, &reply)

	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	log.Printf("%t", reply)
}
