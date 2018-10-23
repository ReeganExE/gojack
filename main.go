package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	handler := NewServer()

	listener, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal("port 8989 already in use")
	}

	fmt.Printf("Listening on %s\n", listener.Addr())
	http.Serve(listener, handler)
}
