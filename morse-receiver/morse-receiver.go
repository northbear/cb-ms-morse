package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World from Receiver!")
	resp, err := http.Get("http://localhost:5000/")
	if err != nil {
		fmt.Println(err)
	}

}
