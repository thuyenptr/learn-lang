package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type messageResponse struct {
	Message string `json:"message"`  
	Data string `json:"data"`
}

type messageRequest struct {
	Name string `json:"name"`
}

type Handler struct {
	Message string
}

func (handler *Handler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_,_ = fmt.Fprint(w, handler.Message)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/home", homeHandler)
	handler := Handler{"Test"}
	http.Handle("/test", &handler)
	fmt.Println("server listening at 8080")
	_ = http.ListenAndServe(":8080", nil)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Home")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil{
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var rq messageRequest

	err = json.Unmarshal(body, &rq)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	response:= messageResponse{Message : "json response", Data: rq.Name}
	data, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	_,_ = fmt.Fprint(w, string(data))
}
