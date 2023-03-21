package main

import (
	"log"
	"net/http"
)

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	request.Write(writer)
	writer.Write([]byte("\nGreetings from coderschool.vn\n"))
}

func main() {
	log.Println("starting server, listening on port 8001")
	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":8001", nil)
}
