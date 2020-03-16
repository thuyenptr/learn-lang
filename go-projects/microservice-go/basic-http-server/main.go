package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port = 8080
	http.HandleFunc("/", helloHandler)
	fmt.Printf("server listening on port %d", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello world")
}