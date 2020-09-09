package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	fmt.Println("Hello, World from Responder!")
	server := &http.Server{
		Addr: ":5000",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = "0.0.0.0"
		}
		fmt.Fprintf(w, "Hello, World from Responder! %q\n", ip)
		log.Printf("%+v\n", r)
	})

	log.Fatal(server.ListenAndServe())
}
