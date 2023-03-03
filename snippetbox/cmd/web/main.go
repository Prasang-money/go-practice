package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", home)
	serverMux.HandleFunc("/snippet", showSnippet)
	serverMux.HandleFunc("/snippet/create", createSnippet)

	con, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	log.Println("starting server on port 8085")
	if err := http.Serve(con, serverMux); err != nil {
		panic(err)
	}

}
