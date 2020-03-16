package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	message string
}

func (handler *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, handler.message); err != nil {
		_ = fmt.Errorf("something wrong")
	}
}


func main() {
	handler := &myHandler{"Hello world"}

	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("error")
	}
}
