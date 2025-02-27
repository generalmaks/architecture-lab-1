package main

import (
	"lab_1/handlers"
	"log"
	"net"
	"net/http"
)

const host = "localhost"
const port = "8795"

func main() {
	address := net.JoinHostPort(host, port)
	for path, handler := range handlers.Handlers {
		controller := func(response http.ResponseWriter, request *http.Request) {
			for _, method := range handler.Methods {
				if method == request.Method {
					handler.Controller(response, request)
					return
				}
				response.WriteHeader(http.StatusMethodNotAllowed)
				message := "Method " + request.Method + " not allowed"
				response.Write([]byte(message))
			}
		}
		http.HandleFunc(path, controller)
	}
	fail := http.ListenAndServe(address, nil)
	log.Fatal(fail)
}
