package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValidationRequest struct {
	Data string `json:"data"`
}


type validationHandler struct {
	next http.Handler
}

type destinationHandler struct {
	Result string `json:"result"`
}

func newValidationHandler(handler http.Handler) http.Handler {
	return validationHandler{handler}
}


func (handler validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request ValidationRequest

	_ = decoder.Decode(&request)

	fmt.Println(request.Data)
	_,_ = fmt.Fprintln(w, request.Data)

	if request.Data != "Nic" {
		_, _ = fmt.Fprintln(w, "Need Nic data")
		return
	}

	handler.next.ServeHTTP(w, r)
}

func (desHandler destinationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, desHandler.Result)
}

func main()  {
	des := destinationHandler{"This is destination"}
	fmt.Println("server listening at 8080")
	_ = http.ListenAndServe(":8080", newValidationHandler(des))
}
