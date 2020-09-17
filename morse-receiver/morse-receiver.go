package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	NUM_TRYS = 10
	PAUSE    = 10
	DEFAULT_SERVICE_ADDRESS = "127.0.0.1:5000"
)

func main() {
	servaddr := os.Getenv("SERVICE_ADDRESS")
	_, _, err := net.SplitHostPort(servaddr)
	if err != nil {
		log.Println("Error: wrong value of SERVICE_ADDRESS envvar")
		servaddr = DEFAULT_SERVICE_ADDRESS
	}
	url := fmt.Sprintf("http://%s/", servaddr)
	log.Println(url)

	http.DefaultClient.Timeout = 5 * time.Second

	for i := 0; i < NUM_TRYS; i++ {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		log.Printf("received: %s", body)
	}
}
