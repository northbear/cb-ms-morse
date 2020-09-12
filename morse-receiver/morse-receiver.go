package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

const (
	NUM_TRYS = 10
	PAUSE    = 10
)

func main() {
	servaddr := os.Getenv("SERVICE_ADDRESS")
	_, _, err := net.SplitHostPort(servaddr)
	if err != nil {
		fmt.Println("Error: wrong value of SERVICE_ADDRESS envvar")
	}
	url := fmt.Sprintf("http://%s/", servaddr)
	fmt.Println(url)

	for i := 0; i < NUM_TRYS; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("received: %s\n", body)
	}
}
