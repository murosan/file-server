package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "3000", "http port")
	directory := flag.String("d", ".", "document root")
	flag.Parse()

	d := http.Dir(*directory)
	s := http.FileServer(d)

	log.Printf("Serving %s\n", d)
	log.Printf("Start listening on port %s\n", *port)
	log.Printf("Local address => http://localhost:%s", *port)

	err := http.ListenAndServe(":"+*port, s)
	if err != nil {
		log.Fatalf("Error on running file server: %v", err)
	}
}
