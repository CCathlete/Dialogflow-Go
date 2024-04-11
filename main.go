package main

import (
	"Golangbot/pkg/httpHandler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler.RequestHandler)
	fmt.Println("Started Listening...")
	http.ListenAndServe(":5000", nil)
}
