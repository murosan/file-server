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

	log.Printf("Serving %s\n", *directory)
	log.Printf("Start listening on port %s\n", *port)
	log.Printf("On local: http://localhost:%s", *port)

	err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*directory)))
	if err != nil {
		log.Fatalf("Error on running file server: %v", err)
	}
}
