package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

const DEFAULT_APP_PORT = "5000"

func MorseEncoder(m string) string {
	return fmt.Sprintf("Fake morse encoding for message: %s", m)
}

func main() {
	fmt.Println("Hello, World from Responder!")
	port := os.Getenv("APP_PORT")
	_, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		log.Printf("Warning: Cannot parse value of envvar APP_PORT: %s", port)
		log.Printf("Use default value: %s", DEFAULT_APP_PORT)
		port = DEFAULT_APP_PORT
	}
	server := &http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = "0.0.0.0"
		}
		fmt.Fprintf(w, "Hello, World from Responder! %q\n", ip)
		log.Printf("%+v\n", r)
	})
	log.Printf("Listen to port: %s", port)
	log.Fatal(server.ListenAndServe())
}
