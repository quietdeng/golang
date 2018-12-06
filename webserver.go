package main

import (
	"book/route"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	HTTP_HOST = "0.0.0.0"
	HTTP_PORT = 8080
)

func main() {
	server := http.Server{
		Addr:           fmt.Sprintf("%s:%d", HTTP_HOST, HTTP_PORT),
		Handler:        &route.RouteHandle{},
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}

	log.Println(fmt.Sprintf("listen: %d", HTTP_PORT))
	log.Fatal(server.ListenAndServe())
}
