package main

import (
	"../constract"
	"fmt"
	"log"
	"net/rpc"
)

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("dialing: ", err)
		return nil
	}
	return client
}

func main() {
	client := CreateClient()
	if client == nil {
		log.Fatal("error")
		return
	}

	var reply constract.HelloWorldResponse

	args := constract.HelloWorldRequest{Name: "Bill"}

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)

	if err != nil {
		log.Fatal("error: ", err)
	}

	fmt.Println(reply.Message)
}
