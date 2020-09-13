package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

const DEFAULT_APP_PORT = "5000"

type MorseEncoderTable map[rune]string

var encoderTable MorseEncoderTable = MorseEncoderTable{
	'a': "*-", 'b': "-***", 'c': "-*-*", 'd': "-**",
	'e': "*", 'f': "**-*", 'g': "--*", 'h': "****",
	'i': "**", 'j': "J*---", 'k': "-*-", 'l': "*-**",
	'm': "--", 'n': "-*", 'o': "---", 'p': "*--*",
	'q': "--*-", 'r': "*-*", 's': "***", 't': "-",
	'u': "**-", 'v': "***-", 'w': "*--", 'x': "-**-",
	'y': "-*--", 'z': "--**",
	'1': "*----", '2': "**---", '3': "***--", '4': "****-",
	'5': "*****", '6': "-****", '7': "--***", '8': "---**",
	'9': "----*", '0': "-----",
	'.': "*-*-*-", ',': "--**--", '?': "**--**", '\'': "*---*",
	'!': "-*-*--", '/': "-**-*", ':': "---***", ';': "-*-*-*",
	'=': "-***-", '+': "*-*-*", '-': "-****-", '_': "**--*-",
	'"': "*-**-*", '@': "*--*-*",
}

func MorseEncoder(w io.Writer, m string) error {
	var encoded = bufio.NewWriter(w)
	for _, c := range m {
		encoded.WriteString(encoderTable[c])
		encoded.WriteString(" ")
	}
	encoded.Flush()
	return nil
}

func main() {
	log.Println("Start Morse Responder... 73.")
	port := os.Getenv("APP_PORT")
	_, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		log.Printf("Incorrect value of APP_PORT environment variable")
		log.Printf("Use default value: %s", DEFAULT_APP_PORT)
		port = DEFAULT_APP_PORT
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request from %s", r.RemoteAddr)
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("Cannot detect remote address. Use default value: 0.0.0.0")
			ip = "0.0.0.0"
		}
		MorseEncoder(w, ip)
	})

	log.Printf("Listen to port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
