package main

import (
	"../constract"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type HelloWorldHandler struct {

}

type httpHandler struct {
	Message string `json:"message"`
}

func (h *HelloWorldHandler)HelloWorld(args *constract.HelloWorldRequest, response *constract.HelloWorldResponse) error {
	response.Message = "Hello, " + args.Name
	return nil
}

func (h *httpHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, h.Message); err != nil {
		_ = fmt.Errorf("something wrong")
	}
}

func startServer() {

	helloWorld := &HelloWorldHandler{}
	hHandler := &httpHandler{"message from http handler"}
	_ = rpc.Register(helloWorld)
	rpc.HandleHTTP()

	fmt.Println("server listening at 8080")
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for {
		conn, _ := listener.Accept()
		fmt.Println("new client connected ", conn)
		go http.Serve(listener, hHandler)
		go rpc.ServeConn(conn)
	}
}


func main() {
	startServer()
}